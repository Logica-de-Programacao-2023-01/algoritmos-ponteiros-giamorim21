package main

import (
	"fmt"
)

//Crie uma função que receba um ponteiro para uma variável
//argumento e modifique o valor da variável dentro da função.
//Certifique-se de que o ponteiro aponte para uma área de
//memória válida e que a memória seja liberada após o uso.

func modificarVariavel(ptr *int) {
	*ptr = *ptr * 2
}

func main() {
	var valor int = 10
	fmt.Println("Valor antigo:", valor)
	fmt.Println("Endereço de memória antigo:", &valor)
	modificarVariavel(&valor)
	fmt.Println("Valor novo:", valor)
	fmt.Println("Endereço de memória novo:", &valor)
}
