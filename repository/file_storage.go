// FileStorageRepository implements CurrencyRepository interface
// It is used to save currency data to a file
package repository

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/razagr/pensionera/domain"
	"github.com/razagr/pensionera/service"
)

type fileStroageRepository struct {
}

// NewFileStorageRepository creates a new FileStorageRepository
//  @return service.CurrencyRepository
func NewFileStorage() service.CurrencyRepository {
	return &fileStroageRepository{}
}

// Save currency data to a file
//  @receiver r
//  @param currency
//  @param avg
//  @return error
func (r *fileStroageRepository) Save(currency domain.Currency, avg float64) error {
	csvfile, err := os.OpenFile("database.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	defer writer.Flush()

	// Write price using %g precision, which stores precision as much as possible
	price := fmt.Sprintf("%g", avg)
	record := [][]string{{fmt.Sprint(currency.Time), currency.Symbol, price}}
	writer.WriteAll(record)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}
