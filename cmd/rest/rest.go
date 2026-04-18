package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rchmachina/rach-fw/configs"
	"github.com/rchmachina/rach-fw/internal/app"
	delivery "github.com/rchmachina/rach-fw/internal/delivery/rest"

	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load()
	conf, err := configs.LoadConfig()
	if err != nil {
		log.Panic(err)
	}

	log.Println("conf port ", conf.Port)

	container := app.NewContainerUser(conf.Db)

	gin := delivery.NewRouter(container).Setup()

	gin.Run(conf.Port)
}
