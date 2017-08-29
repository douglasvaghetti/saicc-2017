package main

import (
	"fmt"
	"net/http"
)

func visitCount(w http.ResponseWriter, r *http.Request) {
	fmt.Sprintln("Olá, eu não sei o seu número de visitante :(")
	// TODO: conte o número de vezes que a página foi carregada
	// TODO: se o visitante for o número 7, mostre uma mensagem, caso contrário, mostra outra
}

func main() {
	http.HandleFunc("/", visitCount)
	http.ListenAndServe("127.0.0.1:50001", nil)
}
