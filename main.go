package main

import (
	"github.com/joho/godotenv"
	"github.com/razagr/pensionera/config"
	"github.com/razagr/pensionera/repository"
	"github.com/razagr/pensionera/service"
)

func main() {

	//Let's load the .env file, you must need to have this file in the root of the project
	godotenv.Load()

	// Configure the environment variables
	Window, Symbols, FinnHubAPIKey := config.NewConfig().Configuration()

	// get our storage repository to store the data in CSV format
	storage := repository.NewFileStorage()

	// create services for all currencies
	var CurrencyServices = map[string]service.CurrencyService{}
	for s := range Symbols {
		CurrencyServices[s] = service.NewService(Window, s, storage)
	}

	// feed dependcy and Run
	repo := repository.NewFinnHubRepository(Window, Symbols, CurrencyServices, FinnHubAPIKey)
	repo.Run()
}
