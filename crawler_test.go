package main

import (
	"testing"
)

func Test_StartWebCrawler(t *testing.T) {
	webGraph := NewGraph()
	domain := "https://devopers.com.br"
	StartWebCrawler(domain, webGraph)
	webGraph.exportDOTFile(domain)
}

func Test_extractDomain(t *testing.T) {

	tests := []struct {
		name string
		url  string
		want string
	}{
		{"basico", "http://devopers.com.br/pasta1", "devopers.com.br"},
		{"nome", "http://facebook.com/joaopedro.hartmannsalomao", "facebook.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractDomain(tt.url); got != tt.want {
				t.Errorf("extractDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
