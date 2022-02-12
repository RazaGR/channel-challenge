package repository

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/razagr/pensionera/domain"
	"github.com/razagr/pensionera/service"
)

// CurrencyJSON is a helper struct to unmarshal the websocket JSON response
type CurrencyJSON struct {
	Data []domain.Currency `json:"data"`
	Type string            `json:"type"`
}

// repo
//  @param window
//  @param symbols
//  @param CurrencyServices
//  @param APIKey
type repo struct {
	window           int
	symbols          map[string]float32
	CurrencyServices map[string]service.CurrencyService
	APIKey           string
}

// Create a new repository
//  @param window
//  @param symbols
//  @param CurrencyServices
//  @param APIKey
//  @return service.WebSocketRepository
func NewFinnHubRepository(
	window int,
	symbols map[string]float32,
	CurrencyServices map[string]service.CurrencyService,
	APIKey string) service.WebSocketRepository {
	return &repo{window, symbols, CurrencyServices, APIKey}
}

// Run socket services
//  @receiver r
//  @return error
func (r *repo) Run() error {
	w, _, err := websocket.DefaultDialer.Dial("wss://ws.finnhub.io?token="+r.APIKey, nil)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	r.subscribe(w)
	r.start(w)
	return nil
}

// subscribe to the websocket topics
//  @receiver r
//  @param w
func (r *repo) subscribe(w *websocket.Conn) {
	for s := range r.symbols {
		msgReceived, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		w.WriteMessage(websocket.TextMessage, msgReceived)
	}
}

// start fetching data from websocket using  interval set in the r.window varilable
//  @receiver r
//  @param w
func (r *repo) start(w *websocket.Conn) {
	for {

		// respone will save the websocket JSON response
		var respone CurrencyJSON
		// TODO: Talk this issue with the author for emptying buffer from ReadJSON
		err := w.ReadJSON(&respone)
		if err != nil {
			// defer wg.Done()
			panic(err)
		}
		if respone.Type == "trade" {

			// TODO: there are duplicate entries with same symbol, Only difference
			// in the map is the Volume value,  we should discuess if we need to process all of them or not

			// existSymbol helps to avoid duplicate entries in the response
			existSymbol := []string{}

			for _, curr := range respone.Data {

				// check if the symbol already exist in the existSymbol slice
				found, _ := r.found(curr.Symbol, existSymbol)

				// if the symbol does not exist in the slice then add it to
				// the existSymbol slice and also perform the add operation
				if !found {
					existSymbol = append(existSymbol, curr.Symbol)
					go func() {
						err := r.AddToChannel(curr)
						if err != nil {
							panic(err)
						}
					}()
				}
			}

		}
	}

}

// AddToChannel sends the currency data to the
//  @receiver r
//  @param currency
//  @return error
func (r *repo) AddToChannel(currency domain.Currency) error {

	err := r.CurrencyServices[currency.Symbol].AddToChannel(currency)
	if err != nil {
		return err
	}
	return nil
}

// found is a helper function to check if the symbol already exist in the map
// this is required to avoid duplicate entries I was receivng in the websocket JSON response
//  @receiver r
//  @param val
//  @param array
//  @return ok
//  @return i
func (r *repo) found(val string, array []string) (ok bool, i int) {
	for i = range array {
		if ok = array[i] == val; ok {
			return
		}
	}
	return
}
