package main

import (
	"log"

	_ "go.uber.org/automaxprocs"

	config "github.com/thnkrn/comet/api/pkg/config"
	di "github.com/thnkrn/comet/api/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
