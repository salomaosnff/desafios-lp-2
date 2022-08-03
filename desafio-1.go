package main

import (
	"fmt"
	"ler"
	"performance"
)

func main() {
	var numero = ler.Numero("Digite um número: ")
	var result = uint64(1)

	performance.Measure(func() {
		result = fatorialIterativo(numero)
	})

	fmt.Println("O fatorial de", numero, " é ", result)
}

func fatorialIterativo(n int) uint64 {
	var result = uint64(1)

	for i := 2; i <= n; i++ {
		result *= uint64(i)
	}

	return result
}

func fatorialRecursivo(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * fatorialRecursivo(n-1)
}
