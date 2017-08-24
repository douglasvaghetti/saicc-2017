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
		canal <- n
	}
}

func consumidor(canal chan int) {
	for {
		lido := <-canal
		fmt.Println("lido:", lido)
	}
}

func main() {
	canal := make(chan int)
	go produtor(canal)
	go consumidor(canal)
	time.Sleep(time.Microsecond * 50)
}

//END OMIT
