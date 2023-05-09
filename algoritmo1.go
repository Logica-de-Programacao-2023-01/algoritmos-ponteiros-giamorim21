package main

import "fmt"

//Escreva uma função que receba um ponteiro para um inteiro
//e um valor inteiro n. A função deve atualizar o valor do
//inteiro para a soma dos n primeiros números naturais.

func atualizarValor(ptr *int) {
	var soma = 0
	for i := 1; i <= *ptr; i++ {
		soma += i
	}
	*ptr = soma
}
func main() {
	n := 10
	var ptr *int = &n
	atualizarValor(ptr)
	fmt.Println(*ptr)
}
