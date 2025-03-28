package main

import (
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"

	"github.com/mahesh-yadav/stats-app/logger"
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

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func main() {

	logger, logFile := logger.SetupLogger()
	defer logFile.Close()

	arguments := os.Args
	if len(arguments) == 1 {
		logger.Fatal("Need one or more command-line arguments")
	}

	values := make([]float64, 0)

	for i := 1; i < len(arguments); i++ {
		num, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		values = append(values, num)
	}

	if len(values) == 0 {
		logger.Println("Generating 10 random values...")
		for i := 0; i < 10; i++ {
			values = append(values, randomFloat(-10, 10))
		}
	}
	sort.Float64s(values)

	logger.Println("Number of values: ", len(values))
	logger.Println("Min: ", values[0])
	logger.Println("Max: ", values[len(values)-1])

	sum := float64(0)
	for _, num := range values {
		sum = sum + num
	}

	meanValue := sum / float64(len(values))
	logger.Printf("Mean: %.5f\n", meanValue)

	var squaredSum float64
	for _, num := range values {
		squaredSum = squaredSum + math.Pow((num-meanValue), 2)
	}
	standardDeviation := math.Sqrt(squaredSum / float64(len(values)))
	logger.Printf("Standard Deviation: %.5f\n", standardDeviation)

	normalized := normalize(values, meanValue, standardDeviation)
	logger.Println("Normalized Values: ", normalized)
}
