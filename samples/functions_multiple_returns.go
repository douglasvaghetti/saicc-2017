package main

import "fmt"

func divisaoEResto(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	quociente, resto := divisaoEResto(5, 2)
	fmt.Println("quociente:", quociente, "resto:", resto)
}
