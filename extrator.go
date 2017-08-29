package main

import (
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func getHref(t html.Token) (bool, string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			return true, a.Val
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return false, ""
}

// globalize retorna a url global Ex: http://domini/asd.txt. se a url de outro domínio retorna ""
func globalize(baseDomain, from, to string) string {
	// se começa com o domínio base, é uma URL válida, apenas retorna ela
	if strings.HasPrefix(to, baseDomain) {
		return to
	}
	// se for uma URL válida, mas não pertence ao domínio, retorna uma string vazia
	if strings.HasPrefix(to, "http") {
		return ""
	}

	// se nao começa
	// query := from[len(baseDomain):] + "/" + to
	// query = filepath.Clean(query) + "/"
	return baseDomain + to
}

func extraiLinks(from, domain string, body io.ReadCloser, urls chan Link) {
	tokenizer := html.NewTokenizer(body)

	for {
		i := tokenizer.Next()
		if i == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "a" {
				ok, to := getHref(token)
				if ok {
					to = globalize(domain, from, to)
					if to != "" {
						var l Link
						l.De = from
						l.Para = to
						// fmt.Println("enviando novo link:", l)
						urls <- l
					}
				}
			}
		} else if i == html.ErrorToken {
			break
		}
	}
	body.Close()
}

// START OMIT
func extrator(dominio string, filaDeExtracao chan string, links chan Link) {
	for url := range filaDeExtracao {
		// fmt.Println("extraindo url:", url)
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		extraiLinks(url, dominio, resp.Body, links)
	}
}

// END OMIT
