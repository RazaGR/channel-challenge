package service

import (
	"github.com/razagr/pensionera/domain"
)

// CurrencyService is used to get the currency data from adapters (PriceProviderRepository)
// and save it to the database (CurrencyRepository)
type CurrencyService interface {

	// GetAverage returns the average price of the currency
	GetAverage() float64
	//AddPrice adds the price to the prices slice and calls GetAverage()
	AddPrice(currency domain.Currency) error
	// AddToChannel adds the price to the channel
	AddToChannel(currency domain.Currency) error
}

// CurrencyRepository is used to save currency data, such as database, file, etc.
type CurrencyRepository interface {

	// Save saves the currency data to the repository
	Save(currency domain.Currency, avg float64) error
}

// DataProviderRepository is used to get currency data, most likely from a websocket,GRPC, or REST API
type PriceProviderRepository interface {
	// Run starts the data provider
	Run() error
}
