package main

import "fmt"

type aluno struct {
	curso        string
	matricula    int
	universidade string
}

func main() {
	var douglas aluno
	douglas.curso = "Engenharia de computação"
	douglas.matricula = 51805
	douglas.universidade = "FURG"

	fmt.Println("minha struct:", douglas)
}
