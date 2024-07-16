package insertdata

import (
	"fmt"
	"os"
)

func LoadCSV(filePath string) ([][]string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file %w", err)
	}
	defer file.Close()

	return nil, nil
}
