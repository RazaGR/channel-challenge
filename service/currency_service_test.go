package service

import (
	"reflect"
	"testing"

	"github.com/razagr/pensionera/domain"
)

func TestNewService(t *testing.T) {
	type args struct {
		window   int
		currency string
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
			if got := NewService(tt.args.window, tt.args.currency); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_AddPrice(t *testing.T) {
	type fields struct {
		currency        string
		window          int
		prices          []float64
		priceSliceIndex int
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
				currency:        tt.fields.currency,
				window:          tt.fields.window,
				prices:          tt.fields.prices,
				priceSliceIndex: tt.fields.priceSliceIndex,
			}
			if err := s.AddPrice(tt.args.currency); (err != nil) != tt.wantErr {
				t.Errorf("service.AddPrice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetAverage(t *testing.T) {
	type fields struct {
		currency        string
		window          int
		prices          []float64
		priceSliceIndex int
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				currency:        tt.fields.currency,
				window:          tt.fields.window,
				prices:          tt.fields.prices,
				priceSliceIndex: tt.fields.priceSliceIndex,
			}
			if got := s.GetAverage(); got != tt.want {
				t.Errorf("service.GetAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initPrices(t *testing.T) {
	type args struct {
		window int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initPrices(tt.args.window); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
