package migrations

import (
	"context"
	"strings"

	"github.com/FindHotel/emspy/internal/app/config"
	"github.com/FindHotel/emspy/pkg/logger"
	"github.com/FindHotel/emspy/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

var (
	migrations = migrate.NewMigrations()
	log        = logger.Must("EMSpy migrations")
)

func DBCommands() *cobra.Command {
	migrations.DiscoverCaller()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Can't initialise config: %s", err)
	}
	var db *bun.DB
	rootCmd := &cobra.Command{
		Use:   "db",
		Short: "manage database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			if db == nil {
				log.Fatal("DB is not configured")
			}
			log.Infow("Start application with configuration", "config", cfg)
		},
	}
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "create migration tables",
		Run: func(cmd *cobra.Command, args []string) {
			if db == nil {
				log.Fatal("DB is not configured")
			}
			migrator := migrate.NewMigrator(db, migrations)
			utils.PanicOnErr(migrator.Init(context.Background()))
		},
	}
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate database",
		Run: func(cmd *cobra.Command, args []string) {
			if db == nil {
				log.Fatal("DB is not configured")
			}
			migrator := migrate.NewMigrator(db, migrations)

			group, err := migrator.Migrate(context.Background())
			if err != nil {
				panic(err)
			}

			if group.ID == 0 {
				log.Info("there are no new migrations to run")
			}

			log.Infof("migrated to %s\n", group)
		},
	}
	rollbackCmd := &cobra.Command{
		Use:   "rollback",
		Short: "rollback the last migration group",
		Run: func(cmd *cobra.Command, args []string) {
			if db == nil {
				log.Fatal("DB is not configured")
			}
			migrator := migrate.NewMigrator(db, migrations)

			group, err := migrator.Rollback(context.Background())
			if err != nil {
				panic(err)
			}

			if group.ID == 0 {
				log.Info("there are no groups to roll back")
			}

			log.Infof("rolled back %s\n", group)
		},
	}
	lockCmd := &cobra.Command{
		Use:   "lock",
		Short: "lock migrations",
		Run: func(cmd *cobra.Command, args []string) {
			if db == nil {
				log.Fatal("DB is not configured")
			}
			migrator := migrate.NewMigrator(db, migrations)

			err := migrator.Lock(context.Background())
			if err != nil {
				panic(err)
			}
		},
	}
	unlockCmd := &cobra.Command{
		Use:   "unlock",
		Short: "unlock migrations",
		Run: func(cmd *cobra.Command, args []string) {
			if db == nil {
				log.Fatal("DB is not configured")
			}
			migrator := migrate.NewMigrator(db, migrations)

			err := migrator.Unlock(context.Background())
			if err != nil {
				panic(err)
			}
		},
	}
	createGoCmd := &cobra.Command{
		Use:   "create_go",
		Short: "create Go migration",
		Run: func(cmd *cobra.Command, args []string) {
			if db == nil {
				log.Fatal("DB is not configured")
			}
			migrator := migrate.NewMigrator(db, migrations)

			name := strings.Join(args, "_")
			mf, err := migrator.CreateGoMigration(context.Background(), name)
			if err != nil {
				panic(err)
			}
			log.Infof("created migration %s (%s)\n", mf.Name, mf.Path)
		},
	}
	createSQLCmd := &cobra.Command{
		Use:   "create_sql",
		Short: "create up and down SQL migrations",
		Run: func(cmd *cobra.Command, args []string) {
			if db == nil {
				log.Fatal("DB is not configured")
			}
			migrator := migrate.NewMigrator(db, migrations)

			name := strings.Join(args, "_")
			files, err := migrator.CreateSQLMigrations(context.Background(), name)
			if err != nil {
				panic(err)
			}

			for _, mf := range files {
				log.Infof("created migration %s (%s)\n", mf.Name, mf.Path)
			}
		},
	}
	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "print migrations status",
		Run: func(cmd *cobra.Command, args []string) {
			if db == nil {
				log.Fatal("DB is not configured")
			}
			migrator := migrate.NewMigrator(db, migrations)

			ms, err := migrator.MigrationsWithStatus(context.Background())
			if err != nil {
				panic(err)
			}
			log.Infof("migrations: %s\n", ms)
			log.Infof("unapplied migrations: %s\n", ms.Unapplied())
			log.Infof("last migration group: %s\n", ms.LastGroup())
		},
	}
	markAppliedCmd := &cobra.Command{
		Use:   "mark_applied",
		Short: "mark migrations as applied without actually running them",
		Run: func(cmd *cobra.Command, args []string) {
			if db == nil {
				log.Fatal("DB is not configured")
			}
			migrator := migrate.NewMigrator(db, migrations)

			group, err := migrator.Migrate(context.Background(), migrate.WithNopMigration())
			if err != nil {
				panic(err)
			}

			if group.ID == 0 {
				log.Infof("there are no new migrations to mark as applied\n")
			}

			log.Infof("marked as applied %s\n", group)
		},
	}

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(rollbackCmd)
	rootCmd.AddCommand(lockCmd)
	rootCmd.AddCommand(unlockCmd)
	rootCmd.AddCommand(createGoCmd)
	rootCmd.AddCommand(createSQLCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(markAppliedCmd)

	return rootCmd
}
