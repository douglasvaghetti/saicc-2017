package main

import (
	"fmt"
	"net/http"
)

// START OMIT
var visitCountN int

func visitCount(w http.ResponseWriter, r *http.Request) {
	visitCountN++
	if visitCountN == 7 {
		fmt.Fprintln(w, "Parabéns!!!111 Você é o visitante de número 7, clique aqui para receber seu prêmio")
	} else {
		fmt.Fprintln(w, "Você é o visitante de número", visitCountN)
	}
}

// END OMIT

func main() {
	http.HandleFunc("/", visitCount)
	fmt.Println("servidor rodando em http://127.0.0.1:50001")
	http.ListenAndServe("127.0.0.1:50001", nil)
}
