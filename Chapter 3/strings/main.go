package main

import (
	"bytes"
	"fmt"
	"strings"
)

// Exercise 3.10
// comma inserts comma in a non-negative decimal integer string
// after every three places. "12345" becomes "12,345"
func comma(s string, out *bytes.Buffer) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	// get the fractional part of dividing n by 3 as it would tell
	// us how many string needs to come first before the comma
	fractionalPart := n % 3
	// get the total number of 3 pairs grouping we can have from the string
	totalThreePairGrouping := n / 3

	// if the fractional part is not zero, then add comma after it
	if fractionalPart != 0 {
		out.WriteString(s[:fractionalPart])
		out.WriteString(",")
	}

	// now we go over the three pair grouping, slice the array in pairs of threes
	// starting from the fractionalPart
	for i := range totalThreePairGrouping {
		// start from the fractional part which has the prefix already, then multiply
		// our current group range i.e group 1, group 2 etc. by 3
		start := fractionalPart + i*3
		// end is nothing but the start plus 3 offset
		end := start + 3
		out.WriteString(s[start:end])
		// to prevent writing a comma at the end of the last group,
		// we need to check if i is the last group, if it is
		// then don't write a comma
		if i != totalThreePairGrouping-1 {
			out.WriteString(",")
		}
	}

	return out.String()
}

// Exercise 3.11
// decimalComma inserts comma in a non-negative decimal string
// after every three places. "12345" becomes "12,345"
// "12345.678" becomes "12,345.678"
func decimalComma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	var buf bytes.Buffer
	var optionalSign string

	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		optionalSign = s[0:1]
		s = s[1:]
	}
	if optionalSign != "" {
		buf.WriteString(optionalSign)
	}

	if periodIndex := strings.LastIndex(s, "."); periodIndex != -1 {
		comma(s[:periodIndex], &buf)
		buf.WriteString(".")
		comma(s[periodIndex+1:], &buf)
	} else {
		comma(s, &buf)
	}
	return buf.String()
}

// Exercise 3.12
// An anagram is a word or phrase made by transposing the letters of another word
// or phrase; The word "secure" is an anagram of "rescue." each letter occurring
// the same amount of time
func isAnagram(firstWord, secondWord string) bool {
	// We solve this by creating a map of the rune of the first word and the total
	// number of occurrence
	firstWordMap := make(map[rune]int)
	for _, v := range firstWord {
		firstWordMap[v]++
	}

	for _, v := range secondWord {
		if _, ok := firstWordMap[v]; ok {
			firstWordMap[v]--
		} else {
			return false
		}
		if firstWordMap[v] == 0 {
			delete(firstWordMap, v)
		}
	}

	return len(firstWordMap) == 0

}

func main() {
	var out bytes.Buffer
	fmt.Println(comma("1234567890", &out))
	fmt.Println(decimalComma("123456789.22345"))
	fmt.Println(isAnagram("secure", "rescue"))
	fmt.Println(isAnagram("iceman", "cinema"))
}
