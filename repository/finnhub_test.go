//Finnhub adapatar implements PriceProviderRepository interface
// it is a source adapter for the Finnhub Websocket API
package repository

import (
	"errors"
	"reflect"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/razagr/pensionera/domain"
	"github.com/razagr/pensionera/service"
	"github.com/stretchr/testify/mock"
)

func TestNewFinnHubRepository(t *testing.T) {
	type args struct {
		symbols map[string]float32
		APIKey  string
	}
	tests := []struct {
		name string
		args args
		want service.PriceProviderRepository
	}{
		{
			name: "FinnhubRepository",
			args: args{
				symbols: map[string]float32{
					"USD": 0.01,
					"EUR": 0.01,
					"GBP": 0.01,
				},
				APIKey: "",
			},
			want: &repo{
				symbols: map[string]float32{
					"USD": 0.01,
					"EUR": 0.01,
					"GBP": 0.01,
				},
				APIKey: "",
			},
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFinnHubRepository(tt.args.symbols, tt.args.APIKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFinnHubRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

type repoMock struct {
	mock.Mock
	symbols map[string]float32
	APIKey  string
}

func (m *repoMock) Run(channels map[string]chan domain.Currency) error {
	args := m.Called()
	return args.Get(0).(error)
}
func (m *repoMock) subscribe(w *websocket.Conn) error {
	args := m.Called()
	return args.Get(0).(error)
}
func (m *repoMock) startListening(w *websocket.Conn, channels map[string]chan domain.Currency) error {
	args := m.Called()
	return args.Get(0).(error)
}
func Test_repo_Run(t *testing.T) {
	type fields struct {
		symbols map[string]float32
		APIKey  string
	}
	type args struct {
		channels map[string]chan domain.Currency
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "Run",
			fields: fields{
				symbols: map[string]float32{
					"USD": 0.01,
					"EUR": 0.01,
					"GBP": 0.01,
				},
				APIKey: "",
			},
			args: args{
				channels: map[string]chan domain.Currency{
					"USD": make(chan domain.Currency),
					"EUR": make(chan domain.Currency),
					"GBP": make(chan domain.Currency),
				},
			},
			want: nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoMock{
				symbols: tt.fields.symbols,
				APIKey:  tt.fields.APIKey,
			}
			r.On("Run", mock.Anything).Return(errors.New("should return an error"))
			if err := r.Run(tt.args.channels); err.Error() != "should return an error" {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func Test_repo_subscribe(t *testing.T) {
	type fields struct {
		symbols map[string]float32
		APIKey  string
	}
	type args struct {
		w *websocket.Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "subscribe",
			fields: fields{
				symbols: map[string]float32{
					"USD": 0.01,
					"EUR": 0.01,
					"GBP": 0.01,
				},
				APIKey: "",
			},
			args: args{
				w: &websocket.Conn{},
			},
			want: nil,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoMock{
				symbols: tt.fields.symbols,
				APIKey:  tt.fields.APIKey,
			}
			r.On("subscribe", mock.Anything).Return(errors.New("should return an error"))
			if err := r.subscribe(tt.args.w); err.Error() != "should return an error" {
				t.Errorf("subscribe() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func Test_repo_startListening(t *testing.T) {
	type fields struct {
		symbols map[string]float32
		APIKey  string
	}
	type args struct {
		w        *websocket.Conn
		channels map[string]chan domain.Currency
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "startListening",
			fields: fields{
				symbols: map[string]float32{
					"USD": 0.01,
					"EUR": 0.01,
					"GBP": 0.01,
				},
				APIKey: "",
			},
			args: args{
				w: &websocket.Conn{},
				channels: map[string]chan domain.Currency{
					"USD": make(chan domain.Currency),
					"EUR": make(chan domain.Currency),
					"GBP": make(chan domain.Currency),
				},
			},
			want: nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repoMock{
				symbols: tt.fields.symbols,
				APIKey:  tt.fields.APIKey,
			}
			r.On("startListening", mock.Anything, mock.Anything).Return(errors.New("should return an error"))
			if err := r.startListening(tt.args.w, tt.args.channels); err.Error() != "should return an error" {
				t.Errorf("startListening() error = %v, wantErr %v", err, tt.want)
			}

		})
	}
}

func Test_repo_duplicate(t *testing.T) {
	type fields struct {
		symbols map[string]float32
		APIKey  string
	}
	type args struct {
		val   string
		array []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOk bool
		wantI  int
	}{
		{
			name: "duplicate",
			fields: fields{
				symbols: map[string]float32{
					"USD": 0.01,
					"EUR": 0.01,
					"GBP": 0.01,
				},
				APIKey: "",
			},
			args: args{
				val:   "USD",
				array: []string{"USD", "EUR", "GBP"},
			},
			wantOk: true,
			wantI:  0,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				symbols: tt.fields.symbols,
				APIKey:  tt.fields.APIKey,
			}
			gotOk, gotI := r.duplicate(tt.args.val, tt.args.array)
			if gotOk != tt.wantOk {
				t.Errorf("repo.duplicate() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if gotI != tt.wantI {
				t.Errorf("repo.duplicate() gotI = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}
