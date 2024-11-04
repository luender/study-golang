package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
    API APIConfig
    DB  DBConfig
}

type APIConfig struct {
    Port string
}

type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    Name     string
}

func init() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    viper.AutomaticEnv()

    viper.SetDefault("API_PORT", "8080")
    viper.SetDefault("DB_HOST", "localhost")
    viper.SetDefault("DB_PORT", "5432")
    viper.SetDefault("DB_USER", "postgres")
    viper.SetDefault("DB_PASSWORD", "postgres")
    viper.SetDefault("DB_NAME", "postgres")
}

func Load() error {
    cfg = new(config)

    cfg.API.Port = viper.GetString("API_PORT")

    cfg.DB.Host = viper.GetString("DB_HOST")
    cfg.DB.Port = viper.GetString("DB_PORT")
    cfg.DB.User = viper.GetString("DB_USER")
    cfg.DB.Password = viper.GetString("DB_PASSWORD")
    cfg.DB.Name = viper.GetString("DB_NAME")

    return nil
}

func GetDB() DBConfig {
    return cfg.DB
}

func GetPort() string {
    return cfg.API.Port
}
