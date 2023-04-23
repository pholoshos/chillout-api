// internal/config/config.go

package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

// Config contains the application configuration settings.
type Config struct {
	Port           string
	DatabaseURL    string
	AccessTokenKey string
	RefreshTokenKey string
	AccessTokenDuration time.Duration
	RefreshTokenDuration time.Duration
}

// NewConfig creates a new instance of Config by loading configuration settings from environment variables.
func NewConfig() (*Config, error) {
	viper.SetDefault("PORT", "8080")
	viper.AutomaticEnv()

	port := viper.GetString("PORT")
	dbURL := viper.GetString("DATABASE_URL")
	accessTokenKey := viper.GetString("ACCESS_TOKEN_KEY")
	refreshTokenKey := viper.GetString("REFRESH_TOKEN_KEY")
	accessTokenDuration := viper.GetDuration("ACCESS_TOKEN_DURATION")
	refreshTokenDuration := viper.GetDuration("REFRESH_TOKEN_DURATION")

	if accessTokenDuration == 0 {
		accessTokenDuration = time.Hour * 24
	}

	if refreshTokenDuration == 0 {
		refreshTokenDuration = time.Hour * 24 * 30
	}

	if port == "" {
		return nil, fmt.Errorf("missing required environment variable PORT")
	}

	if dbURL == "" {
		return nil, fmt.Errorf("missing required environment variable DATABASE_URL")
	}

	if accessTokenKey == "" {
		return nil, fmt.Errorf("missing required environment variable ACCESS_TOKEN_KEY")
	}

	if refreshTokenKey == "" {
		return nil, fmt.Errorf("missing required environment variable REFRESH_TOKEN_KEY")
	}

	return &Config{
		Port:           port,
		DatabaseURL:    dbURL,
			AccessTokenKey: accessTokenKey,
			RefreshTokenKey: refreshTokenKey,
			AccessTokenDuration: accessTokenDuration,
			RefreshTokenDuration: refreshTokenDuration,
	}, nil
}
