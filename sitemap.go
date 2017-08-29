package main

import (
	"fmt"
	"sync"
)

// SiteMap é um grafo onde nodos são urls e as arestas são links
type SiteMap struct {
	links map[string][]string
	mutex sync.Mutex
}

// START OMIT

// Link representa um link entre duas paginas, De e Para são URLs
type Link struct {
	De, Para string
}

// END OMIT

// NewSiteMap inicializa o grafo
func NewSiteMap() *SiteMap {
	var siteMap SiteMap
	siteMap.links = make(map[string][]string)
	return &siteMap
}

// AddLinkCasoNovo adiciona um link ao grafo, retorna true se a página alvo seja nova
func (s *SiteMap) AddLinkCasoNovo(de, para string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	encontrado := false
	for i := 0; i < len(s.links[de]); i++ {
		if para == s.links[de][i] {
			encontrado = true
		}
	}
	if !encontrado {
		s.links[de] = append(s.links[de], para)
	}
	_, paginaNova := s.links[para]
	return !paginaNova
}

// ExportarArquivoDOT faz prints do conteúdo do sitemap em formato .dot,
// para ser visualizado usando graphviz
func (s *SiteMap) ExportarArquivoDOT(dominio string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tamanhoDominio := len(dominio)

	fmt.Println("digraph {")
	for origem, links := range s.links {
		// fmt.Println("node = ", node)
		if len(origem) >= tamanhoDominio {
			// fmt.Println("node = ", node, "limpado:", node[domainLen:])
			fmt.Printf("\"%s\" -> { \"%s\"", origem[tamanhoDominio:], links[0][tamanhoDominio:])
			for _, l := range links[1:] {
				if len(l) > tamanhoDominio {
					fmt.Print(", \"", l[tamanhoDominio:], "\"")
				}
			}
			fmt.Print("}\n\n")
		}
	}
	fmt.Println("}")
}
