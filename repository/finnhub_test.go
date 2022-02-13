//Finnhub adapatar implements PriceProviderRepository interface
// it is a source adapter for the Finnhub Websocket API
package repository

import (
	"reflect"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/razagr/pensionera/domain"
	"github.com/razagr/pensionera/service"
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
					"AAPL": 0.01,
					"MSFT": 0.01,
					"GOOG": 0.01,
				},
				APIKey: "",
			},
			want: &repo{
				symbols: map[string]float32{
					"AAPL": 0.01,
					"MSFT": 0.01,
					"GOOG": 0.01,
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

func Test_repo_Run(t *testing.T) {
	type fields struct {
		symbols map[string]float32
		APIKey  string
	}
	type args struct {
		channels map[string]chan domain.Currency
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
			r := &repo{
				symbols: tt.fields.symbols,
				APIKey:  tt.fields.APIKey,
			}
			if err := r.Run(tt.args.channels); (err != nil) != tt.wantErr {
				t.Errorf("repo.Run() error = %v, wantErr %v", err, tt.wantErr)
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				symbols: tt.fields.symbols,
				APIKey:  tt.fields.APIKey,
			}
			r.subscribe(tt.args.w)
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				symbols: tt.fields.symbols,
				APIKey:  tt.fields.APIKey,
			}
			r.startListening(tt.args.w, tt.args.channels)
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
					"AAPL": 0.01,
					"MSFT": 0.01,
					"GOOG": 0.01,
				},
				APIKey: "",
			},
			args: args{
				val:   "AAPL",
				array: []string{"AAPL", "MSFT", "GOOG"},
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
