package main

import "fmt"

//Escreva uma função que receba um ponteiro
//para um booleano e altere o valor desse
//booleano para o seu inverso.

func inverso(ptr *bool) {
	if *ptr == true {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func main() {
	x := false
	var ptr *bool = &x
	inverso(ptr)
}
