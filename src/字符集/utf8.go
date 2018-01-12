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
	fmt.Println("---------------")
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
	fmt.Println("---------------")
}

func loop() {
	s := "世界"
	
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println("---------------")
}


func main() {
	hello()
	helloCN()
	loop()
}
