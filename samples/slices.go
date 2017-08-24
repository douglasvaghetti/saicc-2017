package main

import "fmt"

func main() {
	var meuSlice []int

	meuSlice = append(meuSlice, 1)
	meuSlice = append(meuSlice, 2)
	meuSlice = append(meuSlice, 3)
	fmt.Println("meuSlice:", meuSlice)

	inicializacaoRapida := []int{4, 5, 6}
	fmt.Println("inicialização rápida:", inicializacaoRapida)
}
