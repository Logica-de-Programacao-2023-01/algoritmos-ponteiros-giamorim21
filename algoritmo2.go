package main

import "fmt"

//Escreva uma função que receba um ponteiro para um inteiro e
//verifique se esse inteiro é par ou ímpar. A função deve atualizar
//o valor do inteiro para 0 se ele for par ou para 1 se for ímpar.

func atualizarNumero(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}
}

func main() {
	x := 10
	var ptr *int = &x
	atualizarNumero(ptr)
	fmt.Println(*ptr)
}
