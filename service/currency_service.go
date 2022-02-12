// CurrencyService is used to get the currency data from adapters (PriceProviderRepository)
// and save it to the database (CurrencyRepository)
package service

import (
	"fmt"
	"math"

	"github.com/razagr/pensionera/domain"
)

var (

	// channels is a map of channels for each currency
	// each country has it's own channel
	channels = map[string]chan domain.Currency{}

	// channelOpened keeps track of which channels are open
	channelOpened = map[string]bool{}
)

// service
type service struct {

	// currency symbol
	currency string

	// window size
	window int

	// prices slice
	prices []float64

	// priceSliceIndex tracks the index of the slice
	priceSliceIndex int

	// storage is the repository for saving currency data
	storage CurrencyRepository
}

// NewService creates a new service for a currency
func NewService(window int, currency string, storage CurrencyRepository) CurrencyService {

	// initialize the prices slice
	prices := make([]float64, window)

	// initialize the priceSliceIndex with 0
	priceSliceIndex := 0
	return &service{currency, window, prices, priceSliceIndex, storage}
}

// AddPrice adds the price to the prices slice and calls GetAverage()
// It also resets the priceSliceIndex to 0 if it reaches the window size
func (s *service) AddPrice(currency domain.Currency) error {
	if math.IsNaN(currency.Price) && math.IsInf(currency.Price, 0) {
		return fmt.Errorf("Invalid price: %f", currency.Price)
	}

	// add the new price to the prices slice
	s.prices[s.priceSliceIndex] = currency.Price

	// our slice is full, lets calculate the moving average and reset the slice
	// we have to subtract 1 from the window because we are using a circular buffer
	if s.priceSliceIndex >= (s.window - 1) {
		avg := s.GetAverage()

		// reset the slice index for circular buffer
		s.priceSliceIndex = 0
		// reset the prices slice values to 0
		s.prices = make([]float64, s.window)

		// starts a goroutine to save the average to the database
		go func() {
			fmt.Printf("-> Currency: %s: Window: %d:  Timestamp: %d: Average: %v\n", s.currency, s.window, currency.Time, avg)
			s.storage.Save(currency, avg)
		}()
	} else {
		// increase the index of the slice
		s.priceSliceIndex++
	}
	fmt.Println(s.priceSliceIndex, "-> ", currency.Symbol, "-> ", currency.Price)
	return nil
}

// Returns the average of the prices slice divided by the window size
func (s *service) GetAverage() float64 {
	var sum float64 = 0

	// count the sum of the prices slice
	for _, price := range s.prices {
		sum += price
	}
	return sum / float64(s.window)
}

// add retrived currency detail from websocket to the map using go channel
func (s *service) AddToChannel(currency domain.Currency) error {

	// check if the channel is open for this currency
	if channelOpened[currency.Symbol] == false {
		// create a new channel for this currency
		channels[currency.Symbol] = make(chan domain.Currency)
		// set the channel as open
		channelOpened[currency.Symbol] = true
		fmt.Println("Channel opened for: ", currency.Symbol)
		// start a goroutine to listen to the channel
		go func(c <-chan domain.Currency) {

			// loop through channel until the channel is closed
			for cur := range c {
				// when channel sends a value then send it to the service for processing
				err := s.AddPrice(cur)
				if err != nil {
					panic(err)
				}
			}
			fmt.Println("Channel closed for: ", currency.Symbol)
		}(channels[currency.Symbol])
	}
	// send the currency to the channel
	channels[currency.Symbol] <- currency
	return nil
}
