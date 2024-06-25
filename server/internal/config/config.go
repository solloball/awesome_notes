package config

import (
    "time"
    "os"
    "log"

    "github.com/ilyakaznacheev/cleanenv"
    "github.com/joho/godotenv"
)

type Config struct {
    Env string `yaml:"env" env-required:"true"`
    StoragePath string `yaml:"sttorage_path" env-required:"true"` 
    HttpServer `yaml:"http_server"`
    Clients ClientsConf `yaml:"clients"`
    AppSecret string `yaml:"app_secret" env-required:"true" env:"APP_SECRE"`
}

type HttpServer struct {
    Address string `yaml:"address" env-default:"localhost:8080"`
    Timeout time.Duration `yaml:"timeout" env-default:"10s"`
    IdleTimeout time.Duration `yaml:"iddle_timeout" env-default:"100s"`
    User string `yaml:"user" env-required:"true"`
    Password string `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

func MustLoad() Config {
    if err := godotenv.Load(".env"); err != nil {
        log.Fatal(err)
    }

    config := os.Getenv("CONFIG_PATH")
    if config == "" {
        log.Fatal("CONFIG_PATH is not set up")
    }


    // check is file exist
    if _, err := os.Stat(config); err != nil {
        log.Fatal(err)
    }

    var conf Config

    if err := cleanenv.ReadConfig(config, &conf); err != nil {
        log.Fatal(err)
    }

    return conf
}

type ClientConf struct {
    Address string `yaml:"address"`
    Timeout time.Duration `yaml:"timeout" env-default:"10s"`
    RetriesCount int `yaml:"retries_count"`
}

type ClientsConf struct {
    SSO ClientConf `yaml:"sso"`
}
