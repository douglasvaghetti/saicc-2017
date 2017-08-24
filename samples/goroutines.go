package main

import (
	"fmt"
	"time"
)

// START OMIT
func repeat10times(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(str)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	go repeat10times("go é demais!")
	go repeat10times("é mesmo!")
	time.Sleep(time.Second * 3)
}

// END OMIT
