package main

import "fmt"

//Escreva uma função swap que receba dois ponteiros
//para int como argumentos e troque
//os valores apontados por eles.

func swap(ptr1 *int, ptr2 *int) {
	var aux = *ptr1
	*ptr1 = *ptr2
	*ptr2 = aux

}

func main() {
	x := 21
	y := 42
	var ptr1 *int = &x
	var ptr2 *int = &y
	fmt.Println("Valor apontado por ptr1:", *ptr2)
	fmt.Println("Valor apontado por ptr2:", *ptr1)
	swap(ptr1, ptr2)
}
