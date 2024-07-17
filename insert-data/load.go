package insertdata

import (
	"encoding/csv"
	"fmt"
	"os"
)

func LoadCSV(filePath string) ([][]string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read CSV file %w", err)
	}

	return records, nil
}
