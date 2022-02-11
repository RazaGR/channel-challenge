package service

import (
	"github.com/razagr/pensionera/domain"
)

// CurrencyService is used to get currency data, most likely from a websocket,GRPC, or REST API
type CurrencyService interface {
	GetAverage() float64
	AddPrice(currency domain.Currency) error
	AddToChannel(currency domain.Currency) error
}

// CurrencyRepository is used to save currency data to a database
type CurrencyRepository interface {
	Save(currency *domain.Currency) error
}

// WebSocketRepository is used to assign rules to a websocket
type WebSocketRepository interface {
	Run() error
}
