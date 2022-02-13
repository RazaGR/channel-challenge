// CurrencyService is used to get the currency data from adapters (PriceProviderRepository)
// and save it to the database (CurrencyRepository)
package service

import (
	"fmt"
	"math"
	"sync"

	"github.com/razagr/pensionera/domain"
)

var (

	// channels is a map of channels for each currency
	// each country has it's own channel
	channels = map[string]chan domain.Currency{}

	// a mutex to prevent race condition on Add()
	addMutex = sync.RWMutex{}
)

// service
type service struct {

	// window size
	window int

	// currency symbol
	symbols map[string]float32

	// prices slice
	prices map[string][]float64

	// priceSliceIndex tracks the index of the slice
	priceSliceIndex map[string]int

	// storage is the repository for saving currency data
	storage CurrencyRepository

	repo PriceProviderRepository
}

// NewService creates a new service for a currency
func NewService(window int, symbols map[string]float32, storage CurrencyRepository, repo PriceProviderRepository) CurrencyService {
	prices := make(map[string][]float64)
	priceSliceIndex := make(map[string]int)

	return &service{window, symbols, prices, priceSliceIndex, storage, repo}
}

func (s *service) Run() {

	// setup the channels
	for c := range s.symbols {
		s.prices[c] = make([]float64, s.window)
		s.priceSliceIndex[c] = 0
		channels[c] = make(chan domain.Currency)
		fmt.Println("Channel opened for: ", c)
		go func(c <-chan domain.Currency) {
			// loop through channel until the channel is closed
			for cur := range c {
				// Lock add operation to save  from race condition
				addMutex.Lock()
				// when channel sends a value then send it to the service for processing
				err := s.addPrice(cur)
				addMutex.Unlock()
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			}
			fmt.Println("Channel closed for: ", c)
		}(channels[c])
	}
	// start the repository to get the currency data
	s.repo.Run(channels)
}

// addPrice adds the price to the prices slice and calls getAverage()
// It also resets the priceSliceIndex to 0 if it reaches the window size
func (s *service) addPrice(currency domain.Currency) error {
	if math.IsNaN(currency.Price) && math.IsInf(currency.Price, 0) {
		return fmt.Errorf("Invalid price: %f", currency.Price)
	}

	// add the new price to the prices slice
	s.prices[currency.Symbol][s.priceSliceIndex[currency.Symbol]] = currency.Price

	// our slice is full, lets calculate the moving average and reset the slice
	// we have to subtract 1 from the window because we are using a circular buffer
	if s.priceSliceIndex[currency.Symbol] >= (s.window - 1) {
		avg := s.getAverage(currency)

		// reset the slice index for circular buffer
		s.priceSliceIndex[currency.Symbol] = 0
		// reset the prices slice values to 0
		s.prices[currency.Symbol] = make([]float64, s.window)

		// starts a goroutine to save the average to the database
		go func() {
			fmt.Printf("-> Currency: %s: window: %d:  Timestamp: %d: Average: %v\n", currency.Symbol, s.window, currency.Time, avg)
			s.storage.Save(currency, avg)
		}()
	} else {
		// increase the index of the slice
		s.priceSliceIndex[currency.Symbol]++
	}

	fmt.Println(s.priceSliceIndex[currency.Symbol], "-> ", currency.Symbol, "-> ", currency.Price)
	return nil
}

// Returns the average of the prices slice divided by the window size
func (s *service) getAverage(currency domain.Currency) float64 {
	var sum float64 = 0

	// count the sum of the prices slice
	for _, price := range s.prices[currency.Symbol] {
		sum += price
	}
	return sum / float64(s.window)
}
