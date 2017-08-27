package main

import "fmt"

func main() {
	meuSlice := []string{"um", "slice", "de", "strings"}
	for indice, valor := range meuSlice {
		fmt.Println("indice:", indice, "valor:", valor)
	}
}
