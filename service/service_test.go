// CurrencyService is used to get the currency data from adapters (PriceProviderRepository)
// and save it to the database (CurrencyRepository)
package service

import (
	"reflect"
	"testing"

	"github.com/razagr/pensionera/domain"
)

func TestNewService(t *testing.T) {
	type args struct {
		window  int
		symbols map[string]float32
		storage CurrencyRepository
		repo    PriceProviderRepository
	}
	tests := []struct {
		name string
		args args
		want CurrencyService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.window, tt.args.symbols, tt.args.storage, tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Run(t *testing.T) {
	type fields struct {
		window          int
		symbols         map[string]float32
		prices          map[string][]float64
		priceSliceIndex map[string]int
		storage         CurrencyRepository
		repo            PriceProviderRepository
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				window:          tt.fields.window,
				symbols:         tt.fields.symbols,
				prices:          tt.fields.prices,
				priceSliceIndex: tt.fields.priceSliceIndex,
				storage:         tt.fields.storage,
				repo:            tt.fields.repo,
			}
			s.Run()
		})
	}
}

func Test_service_addPrice(t *testing.T) {
	type fields struct {
		window          int
		symbols         map[string]float32
		prices          map[string][]float64
		priceSliceIndex map[string]int
		storage         CurrencyRepository
		repo            PriceProviderRepository
	}
	type args struct {
		currency domain.Currency
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				window:          tt.fields.window,
				symbols:         tt.fields.symbols,
				prices:          tt.fields.prices,
				priceSliceIndex: tt.fields.priceSliceIndex,
				storage:         tt.fields.storage,
				repo:            tt.fields.repo,
			}
			if err := s.addPrice(tt.args.currency); (err != nil) != tt.wantErr {
				t.Errorf("service.addPrice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_getAverage(t *testing.T) {
	type fields struct {
		window          int
		symbols         map[string]float32
		prices          map[string][]float64
		priceSliceIndex map[string]int
		storage         CurrencyRepository
		repo            PriceProviderRepository
	}
	type args struct {
		currency domain.Currency
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				window:          tt.fields.window,
				symbols:         tt.fields.symbols,
				prices:          tt.fields.prices,
				priceSliceIndex: tt.fields.priceSliceIndex,
				storage:         tt.fields.storage,
				repo:            tt.fields.repo,
			}
			if got := s.getAverage(tt.args.currency); got != tt.want {
				t.Errorf("service.getAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
