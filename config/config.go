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

// var
//  @param symbols is required to be a map of currency symbols
//  @param window is our window size
//  @param FinnHubAPIKey is the FinnHub API Key
var (
	window        int
	symbols       map[string]float32
	FinnHubAPIKey string
)

func NewConfig() Config {
	if windowEnv := os.Getenv("WINDOWSIZE"); windowEnv != "" {
		var err error
		w, err := strconv.Atoi(windowEnv)
		if err != nil {
			log.Fatal(err)
		}
		window = w
	} else {
		log.Fatal("You must need to provide a WINDOWSIZE environment variable")
	}
	if currencyEnv := os.Getenv("CURRENCY"); currencyEnv != "" {
		sym := make(map[string]float32)
		s := strings.Split(currencyEnv, ",")
		for _, cur := range s {
			sym[cur] = 0
		}
		symbols = sym
	} else {
		log.Fatal("You must need to provide a CURRENCY environment variable")
	}
	if akiKeyEnv := os.Getenv("FINNHUBAPIKEY"); akiKeyEnv != "" {
		FinnHubAPIKey = akiKeyEnv
	} else {
		log.Fatal("You must need to provide a FINNHUBAPIKEY environment variable")
	}
	return &config{window, symbols, FinnHubAPIKey}
}

func (c *config) Configuration() (int, map[string]float32, string) {

	// check if all configuration is correct
	fmt.Printf("Setup is done, You will see result after %d window size \n", window)
	fmt.Println("Window size:", window)
	fmt.Println("Currency: ", symbols)
	fmt.Println("FinnHub API Key:", FinnHubAPIKey)

	return window, symbols, FinnHubAPIKey
}
