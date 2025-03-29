package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadCSVFile(filepath string, delimiter rune) ([][]float64, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = delimiter
	rows, err := reader.ReadAll()
	if err != nil {
		return [][]float64{}, err
	}

	var values [][]float64
	for _, row := range rows {
		var temp []float64

		for _, col := range row {
			num, err := strconv.ParseFloat(strings.TrimSpace(col), 64)
			if err != nil {
				log.Println("Error parsing row: ", err, row)
				break
			}
			temp = append(temp, num)
		}

		if len(temp) == len(row) {
			values = append(values, temp)
		}
	}
	return values, nil
}

func SaveCSVFile(filepath string, data [][]float64, delimiter rune) error {
	csvFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Comma = delimiter

	for _, row := range data {
		temp := make([]string, len(row))
		for i, num := range row {
			temp[i] = strconv.FormatFloat(num, 'f', -1, 64)
		}
		err = csvWriter.Write(temp)
		if err != nil {
			return err
		}
	}
	csvWriter.Flush()
	return nil
}
