package main

import (
	"log"
	"usermanagement/pkg/config"
	dependency "usermanagement/pkg/di"
)

func main() {
	cnf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Unable to load config file due to: ", err)
		return
	}

	server, err := dependency.InitializeDependencies(cnf)

	if err != nil {
		log.Fatal("Unable to get server because of: ", err)
	}

	server.Start(cnf)

}
