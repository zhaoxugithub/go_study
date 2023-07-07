package main

import (
	"fmt"
	"unicode/utf8"
)

/**
 *ClassName strings
 *Description TODO
 *Author 11931
 *Date 2023-07-07:23:45
 *Version 1.0
 **/

func main() {

	s := "yes 我爱慕课网!"
	fmt.Println(s)

	for _, b := range []byte(s) {
		fmt.Println("%X ", b)
	}

	fmt.Println()

	for i, ch := range s {
		fmt.Println("(%d %X) ", i, ch)
	}

	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Println("%c ", ch)
	}

	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
