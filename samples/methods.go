package main

import "fmt"

// START OMIT
type quadrado struct {
	largura, altura int
}

func (q *quadrado) area() int {
	return q.largura * q.altura
}

func main() {
	q := quadrado{4, 5}
	fmt.Println(q.area())
}

// END OMIT
