package main

import (
	"fmt"
	"ler"
	"performance"
)

func main() {
	var numero = ler.Numero("Digite um número: ")

	fmt.Println("Os divisores de ", numero, " são:")

	performance.Measure(func() {
		divisores(numero)
	})
}

func divisores(n int) []int {
	var result = make([]int, 0)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}

	return result
}
