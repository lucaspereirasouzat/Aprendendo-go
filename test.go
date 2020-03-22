package main

import (
	"fmt"
)

// encode para json

func main() {

	// printa o hello world e quebra linha
	fmt.Printf("hello, world\n")
	var a int = lucas(1, 2)
	// printar uma variavel porem ele adiciona o (Extra int =)
	fmt.Printf("a", a)
	// agrupar itens para printar dps
	k := fmt.Sprintln(nil, 1, "lucas")
	fmt.Printf(k)
	j := 125*10 ^ 3
	h := fmt.Sprintln(j)
	// retorno de uma conta
	fmt.Printf(h)
	// printando uma função q retorna uma string
	fmt.Printf(hello())

	//m := Message{"Alice", "Hello", 1294706395881547000}
}

func lucas(a int, b int) int {
	return a + b
}

func add(x int, y int) int {
	return x + y
}

func hello() string {
	return "\n a"
}

// construir um executavel
// go build hello.go

// rodar teste
//go run test.go
