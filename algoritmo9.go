package main

import "fmt"

//Implemente uma função que receba um ponteiro para uma
//struct "Livro" com campos "Título", "Autor" e "Preço",
//e que adicione um desconto de 10% no preço do livro.

type Livro struct {
	titulo string
	autor  string
	preco  float64
}

func main() {
	p := &Livro{
		titulo: "Maus",
		autor:  "LAla",
		preco:  100,
	}
	desconto(p, p.preco)
	fmt.Println(p.preco)
}
func desconto(p *Livro, desconto float64) {
	p.preco = p.preco * 0.9
}
