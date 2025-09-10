// Write a function that counts the number of bits that are different in two SHA256 hashes

package main

import (
	"crypto/sha256"
	"fmt"
)

// use an init function to create the cheat sheet for this
// write the pop count function that sums the total number of ON (1) bits
// in the passed in value.
// now subtract them from each other

var pc [256]byte

// init function creates a table of values from 0 - 255, storing how many ON bits
// each value has

func init() {
	for i := range pc {
		// for index i, it get's the byte value of the quotient of i/2 (this is also same as shifting right in binary or we saying give me the binary of i without the last bit), (i&1) means take i and check if the very last bit is 1. So putting it together this says
		// get me the number of 1s in i/2 (all bits except the last) and then add it with 1 if the last bit in i is ON
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var count int
	for i := range 8 {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

func bitsDiff(b1, b2 *[sha256.Size]byte) int {
	var sum int
	for i := range len(b1) {
		// this is wrong because it's taking the decimal value of the XORed operation and adding it to the sum. Instead of actually counting the bits
		// sum += int(b1[i] ^ b2[i])
		// This is correct since now with pop count we're actually counting the bits that are different
		sum += PopCount(uint64(b1[i] ^ b2[i]))
	}
	return sum
}

func main() {
	c1 := sha256.Sum256([]byte("super secure"))
	c2 := sha256.Sum256([]byte("EXTREMELY SUPER SECURE"))

	fmt.Printf("This is the answer from bitsDiff, given C1: %v and C2 is %v\n the difference is %v\n", c1[0], c2[0], bitsDiff(&c1, &c2))
}
