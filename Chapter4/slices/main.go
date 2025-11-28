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

// Exercise 4.4:Write a version of rotate that operates in a single pass
// To rotate a slice left by n elements, is to apply reverse to the first leading n elements
// next to the remaining elements excluding element 0 to n, then finally reverse the entire
// slice. If you're to rotate right, make the third call first

func rotateLeft(s []int, numRotatePosition int) []int {
	reverse(s[:numRotatePosition])
	reverse(s[numRotatePosition:])
	reverse(s)

	return s
}

func rotatetRight(s []int, numRotatePosition int) []int {
	reverse(s)
	reverse(s[:numRotatePosition])
	reverse(s[numRotatePosition:])

	return s
}

func rotateInSinglePass(s []int, numRotatePosition int) []int {
	var currentPosition int
	for {
		nextPosition := (currentPosition + numRotatePosition) % len(s)
		currentElement, nextElement := s[currentPosition], s[nextPosition]

		s[nextPosition], s[currentPosition] = currentElement, nextElement
		currentElement = nextElement
		currentPosition++

		if nextPosition == 0 {
			break
		}
	}
	return s
}

func main() {
	// intArr := [5]int{1, 2, 3, 4, 5}
	// intSlice := intArr[:]
	// fmt.Printf("this is the array %v\n", intArr)

	// reverseArrPtr(&intArr)
	// fmt.Printf("This is the reversed array using the reverseArrPtr %v\n", intArr)

	// fmt.Printf("this is the slice of the array %v\n", intSlice)
	// reverse(intSlice)
	// fmt.Printf("reversing the slice in place using the original reverse %v\n", intSlice)

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("The slice before rotating %v\n", s)
	// fmt.Printf("The rotated left slice %v\n", rotateLeft(s, 3))
	fmt.Printf("The rotate in single pass slice %v\n", rotateInSinglePass(s, 2))

}
