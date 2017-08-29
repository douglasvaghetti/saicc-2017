package main

import (
	"time"

	_ "golang.org/x/tools/present"
)

// START OMIT

// IniciarWebCrawler inicia o webcrawler
func IniciarWebCrawler(dominio string, siteMap *SiteMap, t time.Duration) {
	links := make(chan Link, 100000)
	filaDeExtracao := make(chan string, 100000)

	for i := 0; i < 100; i++ {
		go consumidor(filaDeExtracao, links, siteMap)
		go extrator(dominio, filaDeExtracao, links)
	}

	filaDeExtracao <- dominio
	time.Sleep(t)
}

func main() {
	siteMap := NewSiteMap()
	dominio := "https://devopers.com.br"
	IniciarWebCrawler(dominio, siteMap, time.Second*10)
	siteMap.ExportarArquivoDOT(dominio)
}

// END OMIT
