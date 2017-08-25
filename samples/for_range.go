package main

import "fmt"

func main() {
	meuSlice := []string{"um", "slice", "de", "strings"}
	for i, v := range meuSlice {
		fmt.Println("i:", i, "v:", v)
	}
}
