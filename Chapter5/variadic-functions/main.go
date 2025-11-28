package main

import "fmt"

func max(vals ...int) int {
	maxNumber := 0
	if len(vals) == 0 {
		return maxNumber
	}
	for _, val := range vals {
		if val > maxNumber {
			maxNumber = val
		}
		continue
	}
	return maxNumber
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	minNumber := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] < minNumber {
			minNumber = vals[i]
		}
		continue
	}
	return minNumber
}

func main() {
	fmt.Println(max(2, 5, 8, 4, 12))         // should return 12
	fmt.Println(max())                       // should return 0
	fmt.Println(min(1, 4, 9, 0, -2, -9, 10)) // should return -9
	fmt.Println(min())                       // should return 0
}
