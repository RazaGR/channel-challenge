//Finnhub adapatar implements PriceProviderRepository interface
// it is a source adapter for the Finnhub Websocket API
package repository

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/razagr/pensionera/domain"
	"github.com/razagr/pensionera/service"
)

// CurrencyJSON saves slice of Currency and its type
type CurrencyJSON struct {
	Data []domain.Currency `json:"data"`
	Type string            `json:"type"`
}

// repo implements the PriceProviderRepository interface
type repo struct {

	// Window size
	Window int

	// currency Symbols
	Symbols map[string]float32

	// // CurrencyService map for each currency symbol
	// CurrencyService service.CurrencyService

	// APIKey is used to authenticate the websocket connection
	APIKey string
}

// NewFinnHubRepository is a constructor for the Finnhub adapter
func NewFinnHubRepository(

	Window int,
	Symbols map[string]float32,
	// CurrencyService service.CurrencyService,
	APIKey string) service.PriceProviderRepository {
	return &repo{Window, Symbols, APIKey}
}

// Run starts the websocket connection and calls the subscribe and start functions
func (r *repo) Run(currencyService service.CurrencyService) error {
	w, _, err := websocket.DefaultDialer.Dial("wss://ws.finnhub.io?token="+r.APIKey, nil)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// subscribe to the websocket
	r.subscribe(w)

	// start listening to the websocket
	r.startListening(w, currencyService)
	return nil
}

// subscribe to the websocket
func (r *repo) subscribe(w *websocket.Conn) {
	for s := range r.Symbols {
		fmt.Println("Subscribing to ", s)
		msgReceived, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		w.WriteMessage(websocket.TextMessage, msgReceived)
	}
}

// start listening to the websocket and passes the pricing data to the CurrencyServices
func (r *repo) startListening(w *websocket.Conn, currencyService service.CurrencyService) {
	for {

		// respone will save the websocket JSON data
		var respone CurrencyJSON
		err := w.ReadJSON(&respone)
		if err != nil {
			panic(err)
		}
		// as per challenge requirement, we only need to process the data if it is a type of trade
		if respone.Type == "trade" {

			// TODO: there are duplicate entries with same symbol, Only difference
			// in the map is the Volume value,  we should discuess if we need to process all of them or not

			// existSymbol helps to avoid duplicate entries in the response
			existSymbol := []string{}

			for _, curr := range respone.Data {

				// check if symbol already exist in the existSymbol slice
				found, _ := r.found(curr.Symbol, existSymbol)

				// if  symbol does not exist in the slice then add it to
				//  existSymbol slice and also perform the add operation
				if !found {
					existSymbol = append(existSymbol, curr.Symbol)

					// start a new goroutine
					go func() {

						// send the pricing data to the CurrencyService
						err := currencyService.AddToChannel(curr)
						if err != nil {
							panic(err)
						}
					}()
				}
			}

		}
	}

}

// found is a function that checks if the symbol already exist in the slice
// this is required to avoid duplicate entries I was receivng in the websocket JSON response
func (r *repo) found(val string, array []string) (ok bool, i int) {
	for i = range array {
		if ok = array[i] == val; ok {
			return
		}
	}
	return
}
