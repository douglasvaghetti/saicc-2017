package main

import (
	"time"

	_ "golang.org/x/tools/present"
)

// Consumer e a goroutine responsável por ver abrir os links extraídos e adiciona-los ao grafo
// se se o o link de destino for novo, ele deve ser adicionado a fila de extração
// Parâmetros:
// toBeExtracted -> um channel de strings, onde serão enviados as páginas cujos links serão extraídos
// links -> um canal de links, é por onde serão retornados os links econtrados pelo extrator
// w -> WebGraph, é o grafo dentro da página, deve ser preenchido com os links.
func consumer(toBeExtracted chan string, links chan Link, w *WebGraph) {
	// TODO Fique lendo o channel de links (dica: tem um for que serve especificamente para ler channels)
	// TODO: De uma olhada nos métodos que o tipo WebGraph implementa (webgraph.go), você vai precisar deles para preencher o grafo
	// TODO: Caso o destino do link inserido for uma página nova, adiciona ela na fila de extração

}

// StartWebCrawler inicia o webcrawler
func StartWebCrawler(domain string, webGraph *WebGraph, t time.Duration) {
	// TODO: crie um channel do tipo Link, use um buffer grande, é lá que os links extraídos serão escritos
	// TODO crie um channel do tipo string, use um buffer grande,  é lá que as URLs que serão extraídas serão escritas
	// TODO: inicie as goroutines de consumidor e extrator
	// TODO: inicie o processo botando o domínio escolhido no canal extrator

	time.Sleep(time.Second * 10) //deixe esse sleep aqui
}

func main() {
	webGraph := NewGraph()
	domain := "http://devopers.com.br"
	StartWebCrawler(domain, webGraph, time.Second*5)
	webGraph.ExportDOTFile(domain)
}
