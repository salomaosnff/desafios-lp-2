package main

import (
	"fmt"
	"ler"
	"performance"
)

func main() {
	var in = uint(ler.Numero("Digite um número: "))

	performance.Measure(func() {
		for L := uint(1); L <= in; L++ {
			fmt.Printf("%d  %d  %d\r\n", L, L*L, L*L*L)
		}
	})
}
