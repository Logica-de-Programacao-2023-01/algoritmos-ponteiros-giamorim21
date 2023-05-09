package main

import "fmt"

//Escreva uma função em Go que receba um ponteiro para uma variável
//inteira e atualize o valor da variável para a soma dos valores dos
//seus dois últimos dígitos (por exemplo, se o valor inicial da
//variável for 1234, o novo valor será 3+4=7).

func atualizarVariavel(ptr *int) {
	valor := *ptr % 100
	soma := (valor / 10) + (valor % 10)
	*ptr = soma
}

func main() {
	var n int = 1234
	atualizarVariavel(&n)
	fmt.Println(n)
}
