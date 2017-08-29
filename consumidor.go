package main

// Consumidor é a goroutine responsável por ver abrir os links extraídos e adiciona-los ao grafo
// se se o o link de destino for novo, ele deve ser adicionado a fila de extração
// Parâmetros:
// filaDeExtracao -> um channel de strings, onde serão enviados as páginas cujos links serão extraídos
// links -> um canal de links, é por onde serão retornados os links econtrados pelo extrator
// siteMap -> SiteMap, é o grafo dentro da página, deve ser preenchido com os links.
// START OMIT
func consumidor(filaDeExtracao chan string, links chan Link, siteMap *SiteMap) {
	for l := range links {
		ehNovo := siteMap.AddLinkCasoNovo(l.De, l.Para)
		if ehNovo {
			filaDeExtracao <- l.Para
		}
	}
}

// END OMIT
