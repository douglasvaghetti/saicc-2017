package main

import "time"

func main() {
	webGraph := NewGraph()
	domain := "http://brrig.com.br"
	StartWebCrawler(domain, webGraph, time.Second*5)
	webGraph.exportDOTFile(domain)
}
