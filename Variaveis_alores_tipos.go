package main

import "fmt"

func main() {
	x := 16
	y := "strings"
	z := true

	//função variadica
	// "_" = blank identifier
	_, erros := fmt.Println("Hello", "OI", 100)

	fmt.Println(erros, x, y, z)
}
