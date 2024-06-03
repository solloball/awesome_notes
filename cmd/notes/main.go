package main


import (
    "fmt"

    "github.com/solloball/aws_tg/internal/config"
)

func main() {
    conf := config.MustLoad()

    fmt.Println(conf)

    //TODO: init logger

    //TODO: init storage

    //TODO: init router

    //TODO: run server
}
