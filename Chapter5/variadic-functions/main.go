package main

import (
	"fmt"
	"log"
)

func max(vals ...int) int {
	var maxNumber int
	if len(vals) == 0 {
		return maxNumber
	}
	maxNumber = vals[0]
	for _, val := range vals[1:] {
		if val > maxNumber {
			maxNumber = val
		}
		continue
	}
	return maxNumber
}

// Max function that requires at least one argument
func max1(vals ...int) (int, error) {
	if len(vals) == 0 {
		log.Fatal("max1 requires at least one argument")
	}
	return max(vals...), nil
}

func min(vals ...int) int {
	var minNumber int
	if len(vals) == 0 {
		return minNumber
	}
	minNumber = vals[0]
	for _, val := range vals[1:] {
		if val < minNumber {
			minNumber = val
		}
		continue
	}
	return minNumber
}

// Min function that requires at least one argument
func min1(vals ...int) (int, error) {
	if len(vals) == 0 {
		log.Fatal("min1 requires at least one argument")
	}
	return min(vals...), nil
}

func main() {
	fmt.Println(max(2, 5, 8, 4, 12))         // should return 12
	fmt.Println(max(-2, -5, -7))             // should return -2
	fmt.Println(max())                       // should return 0
	fmt.Println(min(1, 4, 9, 0, -2, -9, 10)) // should return -9
	fmt.Println(min())                       // should return 0

	fmt.Println(max1(-4, -7, -12))      // should return -4
	fmt.Println(min1(1, 3, 8, 9, 4, 6)) // should return 1
	fmt.Println(max1())                 // should return an error
	fmt.Println(min1())                 // should return an error
}
