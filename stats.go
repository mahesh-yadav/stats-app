package main

import (
	"log"
	"math"
	"os"
	"sort"

	"github.com/mahesh-yadav/stats-app/logger"
	"github.com/mahesh-yadav/stats-app/utils"
)

func normalize(data []float64, mean float64, stdDev float64) []float64 {
	if stdDev == 0 {
		return data
	}

	normalized := make([]float64, len(data))

	for i, val := range data {
		normalized[i] = math.Floor((val-mean)/stdDev*10000) / 10000
	}
	return normalized
}

func MeanAndStdDev(data []float64) (float64, float64) {
	sum := 0.0
	for _, num := range data {
		sum = sum + num
	}

	meanValue := sum / float64(len(data))

	var squared float64
	for _, num := range data {
		squared += math.Pow(num-meanValue, 2)
	}

	stdDev := math.Sqrt(squared / float64(len(data)))
	return meanValue, stdDev
}

func main() {

	logFile := logger.ConfigureLogger()
	defer logFile.Close()

	arguments := os.Args
	if len(arguments) != 3 {
		log.Fatal("Usage: go run main.go <csv filepath> <csv delimiter>\n")
	}

	filePath := arguments[1]
	delimiter := []rune(arguments[2])[0]

	values, err := utils.ReadCSVFile(filePath, delimiter)
	if err != nil {
		log.Fatalf("Error reading CSV file: %v\n", err)
	}

	for i, row := range values {
		sort.Float64s(row)

		mean, stdDev := MeanAndStdDev(row)
		normalizedData := normalize(row, mean, stdDev)

		log.Printf("Row (%d): %v\n", i+1, row)
		log.Println("Min: ", row[0])
		log.Println("Max: ", row[len(row)-1])
		log.Printf("Mean: %.5f\n", mean)
		log.Printf("Standard Deviation: %.5f\n", stdDev)
		log.Println("Normalized Values: ", normalizedData)
		log.Println()
	}
}
