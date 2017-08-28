package main

import (
	"fmt"
	"sync"
)

// WebGraph é um grafo onde nodos são urls e as arestas são links
type WebGraph struct {
	links map[string][]string
	mutex sync.Mutex
}

// Link representa um link entre duas paginas, From e To são URLs
type Link struct {
	From, To string
}

// NewGraph inicializa o grafo
func NewGraph() *WebGraph {
	var grafo WebGraph
	grafo.links = make(map[string][]string)
	return &grafo
}

// AddLinkIfNew adiciona um link ao grafo, retorna true se a página alvo for nova
func (w *WebGraph) AddLinkIfNew(from, to string) bool {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	found := false
	for i := 0; i < len(w.links[from]); i++ {
		if to == w.links[from][i] {
			found = true
		}
	}
	if !found {
		w.links[from] = append(w.links[from], to)
	}
	_, pageSeen := w.links[to]
	return pageSeen
}

func (w *WebGraph) ExportDOTFile(domain string) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	domainLen := len(domain)

	fmt.Println("digraph {")
	for node, links := range w.links {
		if len(node) > domainLen {
			// fmt.Println("node = ", node, "limpado:", node[domainLen:])
			fmt.Printf("\"%s\" -> { \"%s\"", node[domainLen:], links[0][domainLen:])
			for _, l := range links[1:] {
				if len(l) > domainLen {
					fmt.Print(", \"", l[domainLen:], "\"")
				}
			}
			fmt.Println("}\n")
		}
	}
	fmt.Println("}")
}

func panicif(err error) {
	if err != nil {
		panic(err)
	}
}
