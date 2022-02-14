// TODO: These are incomplete/incorrect tests, will come back to them later
package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/razagr/pensionera/domain"
	"github.com/stretchr/testify/mock"
)

type CurrencyRepositoryMock struct {
	mock.Mock
}

func (m *CurrencyRepositoryMock) Save(currency domain.Currency, avg float64) error {
	args := m.Called(currency, avg)
	return args.Error(0)
}

type PriceProviderRepositoryMock struct {
	mock.Mock
}

func (m *PriceProviderRepositoryMock) Run(channels map[string]chan domain.Currency) error {
	args := m.Called()
	// m.On("Run", channels).Return(nil)
	return args.Get(0).(error)
}

type ServiceMock struct {
	mock.Mock
	window          int
	symbols         map[string]float32
	prices          map[string][]float64
	priceSliceIndex map[string]int
	storage         CurrencyRepositoryMock
	repo            PriceProviderRepositoryMock
}

func (m *ServiceMock) Run() error {
	args := m.Called()
	return args.Get(0).(error)
}

func (m *ServiceMock) addPrice(currency domain.Currency) error {

	args := m.Called()
	return args.Get(0).(error)
}
func (m *ServiceMock) getAverage(currency domain.Currency) float64 {
	var sum float64 = 0

	// count the sum of the prices slice
	for _, price := range m.prices["USD"] {
		sum += price
	}
	return sum / float64(m.window)
}

func TestNewService(t *testing.T) {
	storageMock := &CurrencyRepositoryMock{}
	repoMock := &PriceProviderRepositoryMock{}
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
		{
			name: "Test NewService",
			args: args{
				window:  10,
				symbols: map[string]float32{"USD": 1, "EUR": 1.1, "GBP": 1.2},
				storage: storageMock,
				repo:    repoMock,
			},
			want: &service{
				window:          10,
				symbols:         map[string]float32{"USD": 1, "EUR": 1.1, "GBP": 1.2},
				prices:          map[string][]float64{},
				priceSliceIndex: map[string]int{},
				storage:         storageMock,
				repo:            repoMock,
			},
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.window, tt.args.symbols, tt.args.storage, tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() Error \n Is= %v, \n Wa= %v\n", got, tt.want)
			}
		})
	}
}

func Test_service_Run(t *testing.T) {
	storageMock := &CurrencyRepositoryMock{}
	repoMock := &PriceProviderRepositoryMock{}

	type fields struct {
		window          int
		symbols         map[string]float32
		prices          map[string][]float64
		priceSliceIndex map[string]int
		storage         CurrencyRepositoryMock
		repo            PriceProviderRepositoryMock
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			name: "Test Run",
			fields: fields{
				window:          10,
				symbols:         map[string]float32{"USD": 1, "EUR": 1.1, "GBP": 1.2},
				prices:          map[string][]float64{},
				priceSliceIndex: map[string]int{},
				storage:         *storageMock,
				repo:            *repoMock,
			},
			want: nil,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServiceMock{
				window:          tt.fields.window,
				symbols:         tt.fields.symbols,
				prices:          tt.fields.prices,
				priceSliceIndex: tt.fields.priceSliceIndex,
				storage:         tt.fields.storage,
				repo:            tt.fields.repo,
			}
			s.On("Run", mock.Anything).Return(errors.New("should return an error"))
			if err := s.Run(); err.Error() != "should return an error" {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.want)
			}

		})
	}
}

func Test_service_addPrice(t *testing.T) {
	storageMock := &CurrencyRepositoryMock{}
	repoMock := &PriceProviderRepositoryMock{}
	type fields struct {
		window          int
		symbols         map[string]float32
		prices          map[string][]float64
		priceSliceIndex map[string]int
		storage         CurrencyRepositoryMock
		repo            PriceProviderRepositoryMock
	}
	type args struct {
		currency domain.Currency
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Test addPrice",
			fields: fields{
				window:          10,
				symbols:         map[string]float32{"USD": 1, "EUR": 1.1, "GBP": 1.2},
				prices:          map[string][]float64{},
				priceSliceIndex: map[string]int{},
				storage:         *storageMock,
				repo:            *repoMock,
			},
			args: args{
				currency: domain.Currency{
					Symbol: "USD",
					Price:  1.1,
				},
			},
			wantErr: nil,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServiceMock{
				window:          tt.fields.window,
				symbols:         tt.fields.symbols,
				prices:          tt.fields.prices,
				priceSliceIndex: tt.fields.priceSliceIndex,
				storage:         tt.fields.storage,
				repo:            tt.fields.repo,
			}
			s.On("addPrice", mock.Anything).Return(errors.New("should return an error"))
			if err := s.addPrice(tt.args.currency); err.Error() != "should return an error" {
				t.Errorf("addPrice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_getAverage(t *testing.T) {
	storageMock := &CurrencyRepositoryMock{}
	repoMock := &PriceProviderRepositoryMock{}
	type fields struct {
		window          int
		symbols         map[string]float32
		prices          map[string][]float64
		priceSliceIndex map[string]int
		storage         CurrencyRepositoryMock
		repo            PriceProviderRepositoryMock
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
		{
			name: "Test getAverage",
			fields: fields{
				window:          10,
				symbols:         map[string]float32{"USD": 1, "EUR": 1.1, "GBP": 1.2},
				prices:          map[string][]float64{"USD": {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "EUR": {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "GBP": {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
				priceSliceIndex: map[string]int{"USD": 0, "EUR": 0, "GBP": 0},
				storage:         *storageMock,
				repo:            *repoMock,
			},
			args: args{
				currency: domain.Currency{
					Symbol: "USD",
					Price:  1.1,
				},
			},
			want: 5.5,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServiceMock{
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
