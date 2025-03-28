package main

import (
	"math"
	"os"
	"strconv"

	"github.com/mahesh-yadav/stats-app/logger"
)

func main() {

	logger, logFile := logger.SetupLogger()
	defer logFile.Close()

	arguments := os.Args
	if len(arguments) == 1 {
		logger.Fatal("Need one or more command-line arguments")
	}

	var min, max float64
	initialized := false
	nValues := 0
	var sum float64
	validNumbers := make([]float64, 0)

	for i := 1; i < len(arguments); i++ {
		num, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		nValues = nValues + 1
		validNumbers = append(validNumbers, num)
		sum = sum + num
		if !initialized {
			min = num
			max = num
			initialized = true
			continue
		}
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	logger.Println("Number of values: ", nValues)
	logger.Println("Min: ", min)
	logger.Println("Max: ", max)

	if nValues == 0 {
		logger.Println("No valid numbers provided")
		return
	}
	meanValue := sum / float64(nValues)
	logger.Printf("Mean: %.5f\n", meanValue)

	var squaredSum float64
	for _, num := range validNumbers {
		squaredSum = squaredSum + math.Pow((num-meanValue), 2)
	}
	standardDeviation := math.Sqrt(squaredSum / float64(nValues))
	logger.Printf("Standard Deviation: %.5f\n", standardDeviation)
}
