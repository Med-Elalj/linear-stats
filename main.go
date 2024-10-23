package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run stats.go <data_file>")
	}

	data, err := parseFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	n := float64(len(data))
	if n == 0 {
		log.Fatal("No data points found")
	}

	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumX2 := 0.0
	sumY2 := 0.0

	for x, y := range data {
		sumX += float64(x)
		sumY += float64(y)
		sumXY += float64(x) * float64(y)
		sumX2 += float64(x * x)
		sumY2 += float64(y * y)
	}

	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	yIntercept := (sumY - slope*sumX) / n

	pearson := (n*sumXY - sumX*sumY) / (math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY)))

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, yIntercept)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", pearson)
}

func parseFile(filePath string) ([]int, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(file), "\n")

	data := make([]int, 0, len(lines))
	for i, line := range lines {
		if line != "" {
			value, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(i, line)
				return nil, err
			}
			data = append(data, value)
		}
	}

	return data, nil
}
