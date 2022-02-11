package domain

// Currency is our main domain for this serivce
type Currency struct {
	Symbol string  `json:"s"`
	Price  float64 `json:"p"`
	Time   uint64  `json:"t"`
	// Volume    float32 `json:"v"`
	// Condition any     `json:"c"`
}
