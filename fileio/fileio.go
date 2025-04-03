package fileio

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(accountValueFile string) (float64, error) {
	data, err := os.ReadFile(accountValueFile)

	if err != nil {
		return 1000, errors.New("Failed to find a file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

func WriteFloatToFile(accountValueFile string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(accountValueFile, []byte(valueText), 0644)
}
