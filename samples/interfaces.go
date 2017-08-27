package main

import "fmt"

// START OMIT
type geometria interface {
	area() int
}

type quadrado struct {
	largura, altura int
}

func (q quadrado) area() int {
	return q.altura * q.largura
}

func main() {
	var minhaGeometria geometria
	minhaGeometria = quadrado{3, 4}
	fmt.Println(minhaGeometria.area())
}

// END OMIT
