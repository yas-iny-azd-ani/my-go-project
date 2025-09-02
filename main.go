package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

const (
	target    float64 = 1.618
	tolerance float64 = 0.01
)

func isGolden(a, b float64) bool {
	if a < b {
		a, b = b, a
	}
	ratio := (a + b) / a
	distance := math.Abs(ratio - target)
	return distance < tolerance
}

func checkPairs(numbers []float64) {
	startTime := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				if isGolden(numbers[i], numbers[j]) {
					fmt.Printf("%.2f and %.2f have a golden ratio.\n", numbers[i], numbers[j])
				} else {
					fmt.Printf("%.2f and %.2f do not have a golden ratio.\n", numbers[i], numbers[j])
				}
			}(i, j)
		}
	}
	wg.Wait()
	fmt.Println(time.Since(startTime))
}
func checkPairs2(numbers []float64) {
	startTime := time.Now()
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if isGolden(numbers[i], numbers[j]) {
				fmt.Printf("%.2f and %.2f have a golden ratio.\n", numbers[i], numbers[j])
			} else {
				fmt.Printf("%.2f and %.2f do not have a golden ratio.\n", numbers[i], numbers[j])
			}
		}
	}
	fmt.Println(time.Since(startTime))
}

func main() {
	var count int
	fmt.Print("Enter the number of values: ")
	fmt.Scan(&count)
	numbers := make([]float64, count)

	for i := 0; i < count; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scan(&numbers[i])
	}
	checkPairs(numbers)
	checkPairs2(numbers)
}
