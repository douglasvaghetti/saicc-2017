package main

import "fmt"

func addSlices(a, b []int) []int {
	// TODO: criar um slice
	// TODO: fazer com que cada uma das posições
	// dele seja igual as somas de a e b nesta posição
	// os dois slices podem ser de qualquer tamanho, mas o tamanho dos dois é o mesmo
	// para saber o tamanho dos slices, podem usar a função len, Ex: len(a) retorna o tamanho de a
	// Exemplo:
	// a = []int{1,2,3}
	// b = []int{4,5,6}
	// vai retornar []int{5,7,9}
	return nil
}

func main() {
	a := []int{3, 4, 5}
	b := []int{7, 6, 5}
	soma := []int{10, 10, 10}

	if slicesIguais(addSlices(a, b), soma) {
		fmt.Println("Parabéns, sua função está correta!")
	} else {
		fmt.Println("oops, ainda não. Sua função retornou:", addSlices(a, b))
		fmt.Println("para a soma de", a, "e", b)
		fmt.Println("o esperado era", soma)
	}
}

func slicesIguais(a, b []int) bool {
	if a == nil || b == nil || len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
