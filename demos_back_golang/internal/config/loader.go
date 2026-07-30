package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

func Load(configPath string) (Config, error) {
	SetDefault()
	var cfg Config

	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		// Ищем .yaml в стандартных местах
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./configs")
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("failed to read config file: %w", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("failed to unmarshal the config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return cfg, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}

func SetDefault() {
	// App
	viper.SetDefault("app.name", "myapp")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.env", "local")

	// Server
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.read_timeout", "10s")
	viper.SetDefault("server.write_timeout", "10s")
	viper.SetDefault("server.shutdown_timeout", "5s")
	//viper.SetDefault("server.max_header_bytes", 1048576) // 1MB

	// Database
	viper.SetDefault("database.driver", "postgres")
	viper.SetDefault("database.addres", "localhost:5432")
	viper.SetDefault("database.ssl_mode", "disable")
	//viper.SetDefault("database.max_open_conns", 25)
	//viper.SetDefault("database.max_idle_conns", 5)
	//viper.SetDefault("database.conn_max_lifetime", "5m")

}

func (c *Config) Validate() error {
	// Server
	port, err := strconv.Atoi(c.Server.Port)
	if err != nil {
		return fmt.Errorf("cannot converting port from config")
	}
	if port < 1 || port > 65535 {
		return fmt.Errorf("invalid server port: %d", port)
	}

	// Database
	if c.Database.Address == "" {
		return fmt.Errorf("database address is required")
	}
	if c.Database.Username == "" {
		return fmt.Errorf("database username is required")
	}
	if c.Database.Database == "" {
		return fmt.Errorf("database name is required")
	}

	// Environment
	validEnvs := map[string]bool{"local": true, "dev": true, "prod": true}
	if !validEnvs[c.App.Env] {
		return fmt.Errorf("invalid environment: %s", c.App.Env)
	}

	return nil
}
