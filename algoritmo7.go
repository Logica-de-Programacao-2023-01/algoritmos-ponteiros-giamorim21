package main

import "fmt"

//Escreva uma função em Go que receba um ponteiro para um struct
//Conta com campos saldo e titular, e adicione um valor ao saldo
//da conta.

type Conta struct {
	saldo   float64
	titular string
}

func adicionarSaldo(p *Conta, valor float64) {
	p.saldo += valor
}

func main() {
	p := Conta{
		saldo:   100,
		titular: "giovana",
	}
	fmt.Println(p)
	adicionarSaldo(&p, 100)
	fmt.Println(p.saldo)

}
