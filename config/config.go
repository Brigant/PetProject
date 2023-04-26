package config

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Mode string
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	SSLmode  string
}

type Config struct {
	LogLevel string
	// AccessTokenTTL  int
	// RefreshTokenTTL int
	Server          ServerConfig
	DB              PostgresConfig
	Salt            string
	SigningKey      string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

// Allowed logger levels & config key.
const (
	DebugLogLvl = "DEBUG"
	InfoLogLvl  = "INFO"
	ErrorLogLvl = "ERROR"
)

var errNotAllowedLoggelLevel = errors.New("not allowed logger level")

func InitConfig() (Config, error) {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("error while reading config: %w", err)
	}

	loglevel := viper.GetString("loglevel")
	if err := validate(loglevel); err != nil {
		return Config{}, fmt.Errorf("error while cheking allowed loging leveles: %w", err)
	}

	accessTTL := viper.GetInt("access_token_ttl")
	refreshTTL := viper.GetInt("refresh_token_ttl")
	salt := viper.GetString("salt")
	signingKey := viper.GetString("signing_key")

	cfg := Config{
		LogLevel:        loglevel,
		AccessTokenTTL:  time.Duration(accessTTL) * time.Minute,
		RefreshTokenTTL: time.Duration(refreshTTL) * time.Hour,
		Salt:            salt,
		SigningKey:      signingKey,
		Server: ServerConfig{
			Mode: viper.GetString("server.mode"),
			Port: viper.GetString("server.port"),
		},
		DB: PostgresConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Database: viper.GetString("db.name"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			SSLmode:  viper.GetString("db.sslmode"),
		},
	}

	return cfg, nil
}

func validate(logLevel string) error {
	if strings.ToUpper(logLevel) != DebugLogLvl &&
		strings.ToUpper(logLevel) != ErrorLogLvl &&
		strings.ToUpper(logLevel) != InfoLogLvl {
		return errNotAllowedLoggelLevel
	}

	return nil
}
