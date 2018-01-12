package main

import (
    "fmt"
    "strconv"
    //"math"
)

func main(){
    maxUInt8()
    maxInt8()
    binaryOpt()
}

func maxUInt8(){
    var u uint8 = 255
    fmt.Println("\nMax of Unsigned Integer")
    fmt.Println("-----------------------")
    fmt.Printf("u=%d\nu+1=%d\nu*u=%d \n", u, u+1, u*u)
}

func maxInt8(){
    var u int8 = 127
    fmt.Println("\nMax of Integer")
    fmt.Println("-----------------------")
    fmt.Printf("u=%d\nu+1=%d\nu*u=%d \n", u, u+1, u*u)
}

func binaryOpt(){
    var a int64 = 2
    var b int64 = 5
    
    fmt.Println("\nBinary Operations")
    fmt.Println("-----------------------")
    
    fmt.Printf("Integer a=%d (%s)", a, strconv.FormatInt(a, 2))
    fmt.Printf("\nInteger b=%d (%s) \n", a, strconv.FormatInt(b, 2))
    
    fmt.Printf("\nBit AND a&b=%d \n", a&b)
    fmt.Printf("\nBit OR a&b=%d \n", a|b)
    fmt.Printf("\nBit XOR a^b=%d \n", a^b)
    fmt.Printf("\nBit AND NOT a&^b=%d \n", a&^b)
}
