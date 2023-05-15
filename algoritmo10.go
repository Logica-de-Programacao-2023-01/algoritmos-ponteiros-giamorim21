package main

import (
	"fmt"
	"math"
)

func Primo(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func implemento(slicePtr *[]int, size int) {
	primos := make([]int, 0, size)
	num := 2

	for len(primos) < size {
		if Primo(num) {
			primos = append(primos, num)
		}
		num++
	}

	*slicePtr = primos
}

func main() {
	var primos []int
	size := 10

	implemento(&primos, size)

	fmt.Println(primos)
}
