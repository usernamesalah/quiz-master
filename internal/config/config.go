package config

import (
	"github.com/joho/godotenv"

	"github.com/kelseyhightower/envconfig"
)

// Config stores the application configurations.
type Config struct {
	Port      string `envconfig:"PORT" default:"8080"`
	JWTSecret string `envconfig:"JWT_SECRET" default:"secret"`

	Database DatabaseConfig
	Redis    RedisConfig
}

// DatabaseConfig stores database configurations.
type DatabaseConfig struct {
	User     string `envconfig:"DATABASE_USER" required:"true"`
	Host     string `envconfig:"DATABASE_HOST" required:"true"`
	Password string `envconfig:"DATABASE_PASSWORD" required:"true"`
	DB       string `envconfig:"DATABASE_DB" required:"true"`
	URL      string `envconfig:"DATABASE_URL" required:"true"`
	// CONN           string `envconfig:"DATABASE_CONN" required:"true"`
	Driver         string `envconfig:"DATABASE_DRIVER" default:"mysql"`
	MigrationsPath string `envconfig:"DATABASE_MIGRATIONS_PATH" required:"true" default:"file://migrations/mysql"`
}

//RedisConfig stores elastic configurations.
type RedisConfig struct {
	Host     string `envconfig:"REDIS_HOST" default:"elastic" required:"true"`
	Port     string `envconfig:"REDIS_PORT" default:"6729" required:"true"`
	Username string `envconfig:"REDIS_USERNAME"`
	Password string `envconfig:"REDIS_PASSWORD"`
	Db       int    `envconfig:"REDIS_DB"`
	Prefix   string `envconfig:"REDIS_PREFIX"`
	Timeout  string `envconfig:"REDIS_TIMEOUT"`
}

// ReadConfig populates configurations from environment variables.
func ReadConfig() (Config, error) {
	_ = godotenv.Overload()
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
