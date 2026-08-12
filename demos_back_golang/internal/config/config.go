package config

import "fmt"

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Version   string `mapstructure:"version"`
	Env       string `mapstructure:"env"` // local, dev, prod
	JWTSecret string `mapstructure:"jwt_secret"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	//MaxHeaderBytes  int           `mapstructure:"max_header_bytes"`
	//TLS             TLSConfig     `mapstructure:"tls"`
	//CORS            CORSConfig    `mapstructure:"cors"`
}

type DatabaseConfig struct {
	CoonectionString string
	Driver           string `mapstructure:"driver"` // postgres, mysql
	Address          string `mapstructure:"address"`
	Username         string `mapstructure:"username"`
	Password         string `mapstructure:"password"` // Из ENV!
	Database         string `mapstructure:"database"`
	SSLMode          string `mapstructure:"ssl_mode"`
	//MaxOpenConns    int           `mapstructure:"max_open_conns"`
	//MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	//ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

func (c DatabaseConfig) DSN() string {
	// Вид строки подключения "postgres://postgres:postgres@localhost:5432/demos?sslmode=disable"
	return fmt.Sprintf(
		"%s://%s:%s@%s/%s?sslmode=%s",
		c.Driver, c.Username, c.Password, c.Address, c.Database, c.SSLMode,
	)
}
