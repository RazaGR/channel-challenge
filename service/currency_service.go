package service

import (
	"fmt"
	"math"
	"sync"

	"github.com/razagr/pensionera/domain"
)

// var
//  @param wg
//  @param ch
var (
	wg sync.WaitGroup
	ch = make(chan domain.Currency)
)

// service
type service struct {
	currency        string
	window          int
	prices          []float64
	priceSliceIndex int
	storage         CurrencyRepository
}

// Create new service of type CurrencyService
//  @param window
//  @param currency
//  @return CurrencyService
func NewService(window int, currency string, storage CurrencyRepository) CurrencyService {
	prices := initPrices(window)
	priceSliceIndex := 0
	return &service{currency, window, prices, priceSliceIndex, storage}
}

// AddPrice update prices slice with the new price
// It also resets the priceSliceIndex to 0 if it reaches the window size
//  @receiver s
//  @param currency
//  @return error
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

		// reset the slice index
		s.priceSliceIndex = 0
		// reset the prices slice values to 0
		s.prices = initPrices(s.window)
		fmt.Printf("-> Currency: %s: Window: %d:  Timestamp: %d: Average: %v\n", s.currency, s.window, currency.Time, avg)
		s.storage.Save(&currency, avg)
	} else {
		// increase the index of the slice
		s.priceSliceIndex++
	}
	// fmt.Println(s.priceSliceIndex)
	return nil
}

// Returns the average of the prices slice divided by the window size
//  @receiver s
//  @return float64
func (s *service) GetAverage() float64 {
	var sum float64 = 0

	// count the sum of the prices slice
	for _, price := range s.prices {
		sum += price
	}
	return sum / float64(s.window)
}

// add retrived currency detail from websocket to the map using go channel
//  @receiver r
//  @param currency
//  @return error
func (s *service) AddToChannel(currency domain.Currency) error {
	go func(c <-chan domain.Currency) {
		defer wg.Done()

		// loop through channel until the channel is closed
		for cur := range c {
			// when channel sends a value then send it to the service for processing
			err := s.AddPrice(cur)
			if err != nil {
				defer wg.Done()
				panic(err)
			}
		}
		fmt.Println("Channel closed....")
	}(ch)
	ch <- currency
	wg.Wait()
	return nil
}

// This helps to initialze prices slice with 0 values
// we want to keep the system performant by reducing append calls
//  @param window
//  @return []float64
func initPrices(window int) []float64 {

	// create a slice with the size of the window
	prices := make([]float64, window)
	for i := 0; i < window; i++ {
		prices[i] = 0
	}
	return prices
}
