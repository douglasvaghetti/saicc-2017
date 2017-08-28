package main

import (
	"fmt"
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

func extractLinks(from, domain string, body io.ReadCloser, urls chan Link) {
	tokenizer := html.NewTokenizer(body)

	for {
		i := tokenizer.Next()
		if i == html.StartTagToken {
			//fmt.Println("achou uma tag")
			token := tokenizer.Token()
			if token.Data == "a" {
				ok, to := getHref(token)
				// fmt.Println("pegou href", to)
				if ok {
					fmt.Println("globalizando", domain, from, to)
					to = globalize(domain, from, to)
					if to != "" {
						var l Link
						l.From = from
						l.To = to
						urls <- l
						fmt.Println("jogo", l)
					} else {
						fmt.Println("ignorou url:", to, "domain:")
					}
				}
			}
		} else if i == html.ErrorToken {
			break
		}
	}
	body.Close()
}

func extractor(domain string, toBeExtracted chan string, links chan Link) {
	for url := range toBeExtracted {
		resp, err := http.Get(url)
		panicif(err)
		extractLinks(url, domain, resp.Body, links)
	}
}

// func consumer(toBeExtracted chan string, links chan Link, w *WebGraph) {
// 	for l := range links {
// 		isNew := w.AddLinkIfNew(l.From, l.To)
// 		if isNew {
// 			toBeExtracted <- l.To
// 		}
// 	}
// }

// // StartWebCrawler inicia o webcrawler
// func StartWebCrawler(domain string, webGraph *WebGraph, t time.Duration) {
// 	links := make(chan Link, 100000)
// 	toBeExtracted := make(chan string, 100000)

// 	for i := 0; i < 100; i++ {
// 		go consumer(toBeExtracted, links, webGraph)
// 		go extractor(domain, toBeExtracted, links)
// 	}

// 	toBeExtracted <- domain
// 	time.Sleep(t)
// }
