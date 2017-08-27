package main

import (
	"fmt"
	"math/rand"
	"time"
)

//START OMIT
func produtor(canal chan int) {
	for {
		n := rand.Intn(100)
		fmt.Println("vou mandar", n)
		canal <- n
	}
}

func consumidor(canal chan int) {
	for {
		fmt.Println("vou ler algo...")
		lido := <-canal
		fmt.Println("li:", lido)
	}
}

func main() {
	canal := make(chan int, 10)
	go produtor(canal)
	go consumidor(canal)
	time.Sleep(time.Millisecond)
}

//END OMIT
