// Config is used to set global variables
package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config interface {
	Configuration() (int, map[string]float32, string)
}
type config struct {
	window        int
	symbols       map[string]float32
	FinnHubAPIKey string
}

// Setup variables
var (

	// window size
	Window int

	// currency symbols
	Symbols map[string]float32

	// FinnHub API Key
	FinnHubAPIKey string
)

// NewConfig setup all variables
func NewConfig() Config {
	// get the window size
	if windowEnv := os.Getenv("WINDOWSIZE"); windowEnv != "" {
		var err error
		w, err := strconv.Atoi(windowEnv)
		if err != nil {
			log.Fatal(err)
		}
		Window = w
	} else {
		log.Fatal("You must need to provide a WINDOWSIZE environment variable")
	}
	// get the symbols
	if currencyEnv := os.Getenv("CURRENCY"); currencyEnv != "" {
		sym := make(map[string]float32)
		// split the currency symbols to creates its own service
		s := strings.Split(currencyEnv, ",")
		for _, cur := range s {
			sym[cur] = 0
		}
		Symbols = sym
	} else {
		log.Fatal("You must need to provide a CURRENCY environment variable")
	}
	// get the FinnHub API Key
	if akiKeyEnv := os.Getenv("FINNHUBAPIKEY"); akiKeyEnv != "" {
		FinnHubAPIKey = akiKeyEnv
	} else {
		log.Fatal("You must need to provide a FINNHUBAPIKEY environment variable")
	}
	return &config{Window, Symbols, FinnHubAPIKey}
}

// Configuration will return all variables
func (c *config) Configuration() (int, map[string]float32, string) {

	// check if all configuration is correct
	fmt.Printf("Setup is done, You will see result after %d window size \n", Window)
	fmt.Println("Window size:", Window)
	fmt.Println("Currency: ", Symbols)
	fmt.Println("FinnHub API Key:", FinnHubAPIKey)

	return Window, Symbols, FinnHubAPIKey
}
