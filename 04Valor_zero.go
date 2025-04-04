package main

import "fmt"

var a int
var b float64
var c string

func main() {
	fmt.Printf("%v, %T", a, a)
	fmt.Printf("%v, %T", b, b)
	fmt.Printf("%v, %T", c, c)
}
