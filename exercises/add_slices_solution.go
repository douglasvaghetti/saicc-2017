package main

import "fmt"

//START OMIT
func addSlices(a, b []int) []int {
	sum := make([]int, len(a))
	for i := range a {
		sum[i] = a[i] + b[i]
	}
	return sum
}

// END OMIT

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
