package main

import "fmt"

//Escreva uma função em Go que receba um ponteiro para uma string e
//atualize o valor da string para seu reverso.

func atualizarString(ptr *string) {
	slices := []rune(*ptr)
	for i, j := 0, len(slices)-1; i < j; i, j = i+1, j-1 {
		slices[i], slices[j] = slices[j], slices[i]
	}
	*ptr = string(slices)
}

func main() {
	s := "giovana"
	var ptr *string = &s
	atualizarString(ptr)
	fmt.Println(*ptr)
}
