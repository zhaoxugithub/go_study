package main

import "fmt"

/**
 *ClassName nonrepeating
 *Description TODO
 *Author 11931
 *Date 2023-07-07:23:57
 *Version 1.0
 **/

func lengthOfNonRepeatingSubStr(s string) int {

	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
}
