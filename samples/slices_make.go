package main

import "fmt"

func main() {
	sliceMake := make([]int, 5)

	sliceMake[0] = 12
	sliceMake[1] = 34
	sliceMake[2] = 56

	fmt.Println("sliceMake:", sliceMake)

}
