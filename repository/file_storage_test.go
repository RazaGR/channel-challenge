// FileStorageRepository implements CurrencyRepository interface
// It is used to save currency data to a file
package repository

import (
	"reflect"
	"testing"

	"github.com/razagr/pensionera/domain"
	"github.com/razagr/pensionera/service"
)

func TestNewFileStorage(t *testing.T) {
	tests := []struct {
		name string
		want service.CurrencyRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileStorage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileStroageRepository_Save(t *testing.T) {
	type args struct {
		currency domain.Currency
		avg      float64
	}
	tests := []struct {
		name    string
		r       *fileStroageRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &fileStroageRepository{}
			if err := r.Save(tt.args.currency, tt.args.avg); (err != nil) != tt.wantErr {
				t.Errorf("fileStroageRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
