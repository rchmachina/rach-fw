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
	conf := configs.LoadConfig()

	container := app.NewContainerRest(*conf)

	gin := delivery.NewRouter(container).Setup()
	log.Println("running on  ", conf.Port)
	gin.Run(conf.Port)
}
