package csv_process

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type EtheriumPrice struct {
	Date  time.Time
	Price float64
}

func LoadDataFrom(path string) ([]EtheriumPrice, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("error opening prices.csv file: %w", err)
	}

	defer file.Close()

	r := csv.NewReader(file)

	var pricePairs []EtheriumPrice

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Error reading from CSV file: %w", err)
		}

		parsedData, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			return nil, fmt.Errorf("Cannot parse date: %w", err)
		}

		parsedPrice, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, fmt.Errorf("Cannot parse price: %w", err)
		}

		pricePairs = append(pricePairs, EtheriumPrice{
			Price: parsedPrice,
			Date:  parsedData,
		})
	}

	return pricePairs, nil
}
