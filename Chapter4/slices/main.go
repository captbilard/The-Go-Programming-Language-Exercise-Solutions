package main

import "fmt"

// Reverse, this reverses a slice in place

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Exercise 4.3: Rewrite Reverse to use an array pointer instead of a slice

func reverseArrPtr(arrPtr *[5]int) {
	for i, j := 0, len(arrPtr)-1; i < j; i, j = i+1, j-1 {
		arrPtr[i], arrPtr[j] = arrPtr[j], arrPtr[i]
	}
}

func main() {
	intArr := [5]int{1, 2, 3, 4, 5}
	intSlice := intArr[:]
	fmt.Printf("this is the array %v\n", intArr)

	reverseArrPtr(&intArr)
	fmt.Printf("This is the reversed array using the reverseArrPtr %v\n", intArr)

	fmt.Printf("this is the slice of the array %v\n", intSlice)
	reverse(intSlice)
	fmt.Printf("reversing the slice in place using the original reverse %v\n", intSlice)

}
