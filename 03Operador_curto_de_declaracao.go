package main

import "fmt"

var y = "ola bom dia"

func main() {

	// := gopher, serve para declarar variavel
	// gopher ou marmota só pode ser usado em codeblocks, por exemplo, a func main, caso o contrario, declaramos a variavel igual na linha 7.
	// a var y pode ser utilizada em qualquer lugar dentro do package, ja a var x por exemplo, smente dentro da func main.

	x := 10
	z := 5.25
	//y := "ola bom dia"
	fmt.Printf("x: %v - %T\n", x, x)
	fmt.Printf("z: %v - %T\n", z, z)
	fmt.Printf("y: %v - %T\n\n", y, y)

	// além de atribuir um valor, declara outra variavel
	// operando x+20
	x, i := x+20, 30

	fmt.Printf("x: %v - %T\n", x, x)
	fmt.Printf("i: %v - %T\n\n", i, i)

	//bool

	a := 10 == 10
	b := 10 < 20

	fmt.Println(a)
	fmt.Println(b)
}
