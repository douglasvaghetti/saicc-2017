package main

import "fmt"
import "math"

type geometria interface {
	area() float32
}

type quadrado struct {
	largura, altura float32
}

func (q quadrado) area() float32 {
	return q.altura * q.largura
}

// START OMIT
type circulo struct {
	raio float32
}

func (c circulo) area() float32 {
	return math.Pi * c.raio * c.raio
}

// END OMIT

func main() {
	var minhaGeometria geometria
	minhaGeometria = quadrado{3, 3}
	fmt.Println(minhaGeometria.area())
	var outraGeometria geometria
	outraGeometria = circulo{3}
	fmt.Println(outraGeometria.area())
}
