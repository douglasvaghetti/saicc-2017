package saicc

import (
	"io"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func extractLinks(from string, body io.ReadCloser, urls chan Link, w *WebGraph) {
	tokenizer := html.NewTokenizer(body)

	//fmt.Println("criou token")
	domain := extractDomain(from)
	for {
		i := tokenizer.Next()
		if i == html.StartTagToken {
			//fmt.Println("achou uma tag")
			token := tokenizer.Token()
			if token.Data == "a" {
				ok, to := getHref(token)
				//fmt.Println("pegou href")
				if ok {
					if extractDomain(to) == domain {
						//realiza o map
						//fmt.Println("from = ", from, " to = ", to, " ", w.LinkExists(from, to))
						if !w.LinkExists(from, to) {
							w.AddLink(from, to)
							var l Link
							l.From = from
							l.To = to
							urls <- l
						}
					}
				}
			}
		} else if i == html.ErrorToken {
			break
		}
	}
	body.Close()
}
func extractDomain(url string) string {
	pieces := strings.Split(url, "/")
	domain := pieces[2]
	return domain
}

func simpleCrawler(URL string, grafo *WebGraph, links chan Link, waiter *sync.WaitGroup) {
	//preenche o grafo ;
	//1: - pega html do dominio
	//2: - procura um url na pagina
	//3: - dispara um crawler com o novo url
	waiter.Add(1)
	resp, err := http.Get(URL)
	panicif(err)
	extractLinks(URL, resp.Body, links, grafo)
	waiter.Done()
}

func getHref(t html.Token) (bool, string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			if strings.HasPrefix(a.Val, "http") {
				return true, a.Val
			}

		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return false, ""
}
