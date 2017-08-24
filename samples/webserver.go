package main

import (
	"fmt"
	"net/http"
)

// START OMIT
func sayHello(w http.ResponseWriter, r *http.Request) {
	html := `<html>
				<body>
					<h1><center>Olá, mundo</center></h1>
				</body>
			</html>`
	fmt.Fprintln(w, html)
}

func main() {
	http.HandleFunc("/", sayHello)
	fmt.Println("seu servidor está rodando!\nAcesse http://127.0.0.1:50000")
	http.ListenAndServe("127.0.0.1:50000", nil)
}

// END OMIT
