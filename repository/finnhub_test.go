package repository

import (
	"reflect"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/razagr/pensionera/service"
)

func TestNewFinnHubRepository(t *testing.T) {
	type args struct {
		window           int
		symbols          map[string]float32
		CurrencyServices map[string]service.CurrencyService
		APIKey           string
	}
	tests := []struct {
		name string
		args args
		want service.WebSocketRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFinnHubRepository(tt.args.window, tt.args.symbols, tt.args.CurrencyServices, tt.args.APIKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFinnHubRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_Run(t *testing.T) {
	type fields struct {
		window           int
		symbols          map[string]float32
		CurrencyServices map[string]service.CurrencyService
		APIKey           string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				window:           tt.fields.window,
				symbols:          tt.fields.symbols,
				CurrencyServices: tt.fields.CurrencyServices,
				APIKey:           tt.fields.APIKey,
			}
			if err := r.Run(); (err != nil) != tt.wantErr {
				t.Errorf("repo.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repo_subscribe(t *testing.T) {
	type fields struct {
		window           int
		symbols          map[string]float32
		CurrencyServices map[string]service.CurrencyService
		APIKey           string
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
		t.Run(tt.name, func(_ *testing.T) {
			r := &repo{
				window:           tt.fields.window,
				symbols:          tt.fields.symbols,
				CurrencyServices: tt.fields.CurrencyServices,
				APIKey:           tt.fields.APIKey,
			}
			r.subscribe(tt.args.w)
		})
	}
}

func Test_repo_start(t *testing.T) {
	type fields struct {
		window           int
		symbols          map[string]float32
		CurrencyServices map[string]service.CurrencyService
		APIKey           string
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
		t.Run(tt.name, func(_ *testing.T) {
			r := &repo{
				window:           tt.fields.window,
				symbols:          tt.fields.symbols,
				CurrencyServices: tt.fields.CurrencyServices,
				APIKey:           tt.fields.APIKey,
			}
			r.start(tt.args.w)
		})
	}
}

func Test_repo_found(t *testing.T) {
	type fields struct {
		window           int
		symbols          map[string]float32
		CurrencyServices map[string]service.CurrencyService
		APIKey           string
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				window:           tt.fields.window,
				symbols:          tt.fields.symbols,
				CurrencyServices: tt.fields.CurrencyServices,
				APIKey:           tt.fields.APIKey,
			}
			gotOk, gotI := r.found(tt.args.val, tt.args.array)
			if gotOk != tt.wantOk {
				t.Errorf("repo.found() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if gotI != tt.wantI {
				t.Errorf("repo.found() gotI = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}
