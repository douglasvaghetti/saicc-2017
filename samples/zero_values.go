package main

import "fmt"

type aluno struct {
	curso        string
	matricula    int
	universidade string
}

func main() {
	var z aluno

	fmt.Printf("curso:(%s) matricula:(%d) universidade:(%s)",
		z.curso, z.matricula, z.universidade)
}
