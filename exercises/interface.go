package main

import "fmt"

// START OMIT
type geometria interface {
	area() float32
}

type quadrado struct {
	largura, altura float32
}

func (q quadrado) area() float32 {
	return q.altura * q.largura
}

type circulo struct {
	//TODO: quais sao os campos do circulo?
}

// TODO: implementar o método que falta para que o circulo implemente a interface geometria

func main() {
	var minhaGeometria geometria
	minhaGeometria = quadrado{3, 3}
	fmt.Println(minhaGeometria.area())
	var outraGeometria geometria
	outraGeometria = circulo{}
	fmt.Println(outraGeometria.area())
}

// END OMIT
