// Use var para declarar 3 variaveis, utilizando o package level scope. Não atribua valores a variaveis
// Variaveis x, y e z, int, string e bool

package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x = 42
	y = "José"
	z = true

	fmt.Printf("%v\n%v\n%v\n", x, y, z)
}
