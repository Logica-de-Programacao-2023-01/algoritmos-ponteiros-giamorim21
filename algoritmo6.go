package main

import "fmt"

//Escreva uma função em Go que receba um ponteiro para um struct
//Livro com campos título e autor, e altere o título do livro para
//"Desconhecido" se o autor for "Anônimo".

type Book struct {
	titulo string
	autor  string
}

func alterarBook(p *Book) {
	if p.autor == "Anonimo" {
		p.titulo = "Desconhecido"
	}
}
func main() {
	livro1 := Book{
		titulo: "Maus",
		autor:  " Art Spiegelman",
	}
	alterarBook(&livro1)
	fmt.Println(livro1.titulo)

	livro2 := Book{
		titulo: "é assim que acaba",
		autor:  "Anonimo",
	}
	alterarBook(&livro2)
	fmt.Println(livro2.titulo)
}
