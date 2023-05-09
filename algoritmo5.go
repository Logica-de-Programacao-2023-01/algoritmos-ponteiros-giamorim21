package main

import (
	"fmt"
	"math"
)

//Escreva uma função em Go que receba um ponteiro para uma variável
//float64 e atualize o valor da variável para a média aritmética
//entre o seu valor atual e o valor da constante Pi.

func media(ptr *float64) {
	*ptr = (*ptr + math.Pi) / 2.0
}

func main() {
	var n float64 = 3.14
	media(&n)
	fmt.Println(n)
}
