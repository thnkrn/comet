package main

import (
	"log"

	config "github.com/thnkrn/comet/puller/pkg/config"
	di "github.com/thnkrn/comet/puller/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	app, diErr := di.InitializeApp(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		// Note: execute the scheduler inside a goroutine, since the main goroutine will block it
		go app.Scheduler.Start()

		app.Server.Start()
	}
}
