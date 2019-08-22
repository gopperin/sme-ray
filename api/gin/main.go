package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/broker/nsq"
	_ "github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-web"
)

func main() {
	runService()
}

func runService() {
	// Create service
	service := web.NewService(
		web.Name("snc.gc.api.api"),
	)

	service.Init()

	// setup Server Client
	setupVersionClient()
	setupAuthClient()
	setupEventClient()

	// Create RESTful handler (using Gin)

	router := gin.Default()

	// setup router
	SetupVersionRouter(router)
	SetupAuthRouter(router)
	SetupEventRouter(router)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
