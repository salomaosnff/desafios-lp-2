package ler

import "fmt"

func Numero(text string) int {
	var input int

	fmt.Print(text)
	fmt.Scanf("%d", &input)

	return input
}
