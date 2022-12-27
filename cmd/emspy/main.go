package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FindHotel/emspy/cmd/emspy/migrations"
	"github.com/FindHotel/emspy/internal/app/config"
	"github.com/FindHotel/emspy/internal/app/server"
	"github.com/FindHotel/emspy/internal/app/store"
	"github.com/FindHotel/emspy/internal/app/store/file"
	"github.com/FindHotel/emspy/internal/app/store/kinesis"
	"github.com/FindHotel/emspy/pkg/logger"
	"github.com/FindHotel/emspy/pkg/utils"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

func init() {
	rootCmd.PersistentFlags().String("config", "deployment/emspy/config/stg.json", "config file")

	utils.PanicOnErr(
		viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config")),
	)

	rootCmd.AddCommand(migrations.DBCommands())
}

var (
	log     = logger.Must("EMSpy")
	rootCmd = &cobra.Command{
		Use: "content",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.Load()
			if err != nil {
				log.Fatalf("Can't initialise config: %s", err)
			}

			log.Infow("Start application with configuration", "config", cfg)

			if err := mainWithError(cfg); err != nil {
				log.Fatalf("%s", err)
			}
		},
	}
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func mainWithError(cfg *config.Config) error {
	stores := make([]store.Store, 0)

	if cfg.FileStoreConfig != nil {
		path, err := cfg.FileStoreConfig.FullPath()
		if err != nil {
			return err
		}
		fileStore, err := file.New(path)
		if err != nil {
			return err
		}

		if fileStore != nil {
			stores = append(stores, fileStore)
		}
	}

	if cfg.KinesisStoreConfig != nil {
		log.Info("Conecting to the kinesis store...")
		s := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		kinesisStore, err := kinesis.New(s, *cfg.KinesisStoreConfig.StreamName, log)
		if err != nil {
			log.Fatalf("Can't start server because store unavailable: %s", err)
		}

		if kinesisStore != nil {
			stores = append(stores, kinesisStore)
		}
	}

	server := server.New(":8080", stores, log)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return server.Run(ctx)
	})
	g.Go(func() error {
		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal, 1)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Info("Shutting down server...")

		// The context is used to inform the server it has 5 seconds to finish
		// the request it is currently handling
		subCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := server.Shutdown(subCtx); err != nil {
			log.Fatal("Server forced to shutdown: ", err)
		}

		log.Info("Server exiting")
		return nil
	})

	return g.Wait()
}
