package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/razagr/pensionera/repository"
	"github.com/razagr/pensionera/service"
)

// If you need this to run without docker, you can add the values here
var (
	symbols        map[string]float32
	window         int
	FinnHub_APIKey string
)

func main() {

	// Get the environment variables
	if windowEnv := os.Getenv("WINDOWSIZE"); windowEnv != "" {
		var err error
		window, err = strconv.Atoi(windowEnv)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("You must need to provide a WINDOWSIZE environment variable")
		os.Exit(1)
	}
	if akiKeyEnv := os.Getenv("FINNHUBAPIKEY"); akiKeyEnv != "" {
		FinnHub_APIKey = akiKeyEnv
	} else {
		fmt.Println("You must need to provide a FINNHUBAPIKEY environment variable")
		os.Exit(1)
	}
	if currencyEnv := os.Getenv("CURRENCY"); currencyEnv != "" {
		symbols = make(map[string]float32)
		s := strings.Split(currencyEnv, ",")
		for _, cur := range s {
			symbols[cur] = 0
		}
	} else {
		fmt.Println("You must need to provide a CURRENCY environment variable")
		os.Exit(1)
	}

	fmt.Printf("Setup is done, You will see result after %d window size \n", window)
	fmt.Println("Window size:", window)
	fmt.Println("Currency: ", symbols)
	fmt.Println("FinnHub API Key:", FinnHub_APIKey)

	var CurrencyServices = map[string]service.CurrencyService{}
	for s := range symbols {
		CurrencyServices[s] = service.NewService(window, s)
	}
	repository.NewFinnHubRepository(window, symbols, CurrencyServices, FinnHub_APIKey).Run()
}
