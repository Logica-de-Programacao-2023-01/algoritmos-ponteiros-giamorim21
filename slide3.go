package main

import "fmt"

//Escreva uma função que receba um ponteiro para uma
//struct que contém informações
//de um produto (nome, preço e quantidade em estoque).
//A função deve atualizar o preço
//desse produto para um novo valor fornecido como argumento.

type Produto struct {
	nome              string
	preco             int
	quantidadeEstoque int
}

func main() {
	p := &Produto{
		nome:              "garrafa",
		preco:             10,
		quantidadeEstoque: 200,
	}
	mudancaPreco(p, 20)
	fmt.Println(p.preco)
}

func mudancaPreco(p *Produto, novoPreco int) {
	p.preco = novoPreco
}
