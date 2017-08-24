package main

import "fmt"

type aluno struct {
	curso        string
	matricula    int
	universidade string
}

func main() {
	douglas := aluno{
		curso:        "Engenharia de computação",
		matricula:    51805,
		universidade: "FURG",
	}

	fmt.Println("minha struct:", douglas)
}
