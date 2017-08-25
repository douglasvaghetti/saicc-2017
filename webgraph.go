package saicc

import (
	"sync"
)

// WebGraph é um grafo onde nodos são urls e as arestas são links
type WebGraph struct {
	links map[string][]string
	mutex sync.Mutex
}

type Link struct {
	From, To string
}

func NewGraph() *WebGraph {
	var grafo WebGraph
	grafo.links = make(map[string][]string)
	return &grafo
}

// AddLink é uma função que adiciona um link a um WebGraph
func (w *WebGraph) AddLink(from, to string) {
	w.mutex.Lock()
	w.links[from] = append(w.links[from], to)
	w.mutex.Unlock()
}

// LinkExists verifica se existe um link
func (w *WebGraph) LinkExists(from, to string) bool {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	for i := 0; i < len(w.links[from]); i++ {
		if to == w.links[from][i] {
			return true
		}
	}
	return false
}

func panicif(err error) {
	if err != nil {
		panic(err)
	}
}
