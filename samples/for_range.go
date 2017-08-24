package main

import "fmt"

func main() {
	meuSlice := []string{"um", "slice", "de", "strings"}
	for _, v := range meuSlice {
		fmt.Println("leu:", v)
	}
}
