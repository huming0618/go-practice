package main

import (
	"fmt"
	"unicode/utf8"
)

func hello() {
	s := "Hello,Word"
	fmt.Println(len(s))
	fmt.Println(s[0], s[2])
	fmt.Println(s[:])
	fmt.Println(s[:5])
}

func helloCN() {
	s := "世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}

func main() {
	hello()
	helloCN()
}
