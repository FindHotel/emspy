package config

import (
	"errors"
	"path"

	"github.com/spf13/viper"
)

type Config struct {
	Environment *string `json:"environment,omitempty"`

	PostgresqlStoreConfig *PostgresqlStoreConfig `json:"postgresql,omitempty"`
	FileStoreConfig       *FileStoreConfig       `json:"file,omitempty"`
	KinesisStoreConfig    *KinesisStoreConfig    `json:"kinesis,omitempty"`
}

type PostgresqlStoreConfig struct {
	DSN *string `json:"dsn"`
}

type KinesisStoreConfig struct {
	StreamName *string `json:"stream,omitempty"`
}

type FileStoreConfig struct {
	BasePath *string `json:"base"`
	FileName *string `json:"filename"`
}

func (s *FileStoreConfig) FullPath() (string, error) {
	if s.FileName == nil {
		return "", errors.New("filename must be specified")
	}

	if s.BasePath != nil {
		return path.Join(*s.BasePath, *s.FileName), nil
	}

	return *s.FileName, nil
}

func Load() (*Config, error) {
	if err := viper.BindEnv("environment", "STAGE"); err != nil {
		return nil, err
	}

	viper.SetConfigFile(viper.GetString("config"))
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = viper.Unmarshal(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
