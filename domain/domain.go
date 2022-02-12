// Currency domain
package domain

// Currency
type Currency struct {

	// Currency symbol
	Symbol string `json:"s"`

	// Price
	Price float64 `json:"p"`

	// Timestamp
	Time uint64 `json:"t"`
	// Volume    float32 `json:"v"`
	// Condition any     `json:"c"`
}
