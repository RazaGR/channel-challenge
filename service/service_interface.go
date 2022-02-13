package service

import (
	"github.com/razagr/pensionera/domain"
)

// CurrencyService is used to get the currency data from adapters (PriceProviderRepository)
// and save it to the database (CurrencyRepository)
type CurrencyService interface {

	// getAverage returns the average price of the currency
	getAverage(currency domain.Currency) float64
	//addPrice adds the price to the prices slice and calls getAverage()
	addPrice(currency domain.Currency) error
	// Run starts the service
	Run()
}

// CurrencyRepository is used to save currency data, such as database, file, etc.
type CurrencyRepository interface {

	// Save saves the currency data to the repository
	Save(currency domain.Currency, avg float64) error
}

// DataProviderRepository is used to get currency data, most likely from a websocket,GRPC, or REST API
type PriceProviderRepository interface {
	// Run is receving channels to forward the currency data to its channel
	Run(channels map[string]chan domain.Currency) error
}
