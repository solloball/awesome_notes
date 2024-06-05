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
}

type HttpServer struct {
    Address string `yaml:"address" env-default:"localhost:8080"`
    Timeout time.Duration `yaml:"timeout" env-default:"10s"`
    IdleTimeout time.Duration `yaml:"iddle_timeout" env-default:"100s"`
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
