package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/razagr/pensionera/config"
	"github.com/razagr/pensionera/repository"
	"github.com/razagr/pensionera/service"
)

func main() {
	//Let's load the .env file, you must need to have this file in the root of the project
	godotenv.Load()
	// Configure the environment variables
	window, symbols, FinnHubAPIKey := config.NewConfig().Configuration()

	// get our storage repository to store the data in CSV format
	storage := repository.NewFileStorage()

	// check if all configuration is correct
	fmt.Printf("Setup is done, You will see result after %d window size \n", window)
	fmt.Println("Window size:", window)
	fmt.Println("Currency: ", symbols)
	fmt.Println("FinnHub API Key:", FinnHubAPIKey)

	// create services for all currencies
	var CurrencyServices = map[string]service.CurrencyService{}
	for s := range symbols {
		CurrencyServices[s] = service.NewService(window, s, storage)
	}
	// lets feed dependcy injection with all configuration and services and Run
	repository.NewFinnHubRepository(window, symbols, CurrencyServices, FinnHubAPIKey).Run()
}
