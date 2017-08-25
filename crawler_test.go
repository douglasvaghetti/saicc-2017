package saicc

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_crawler(t *testing.T) {
	links := make(chan Link)
	wg := &sync.WaitGroup{}
	graph := NewGraph()
	go simpleCrawler("https://www.ahnegao.com.br/", graph, links, wg)
	go func() {
		time.Sleep(time.Second)
		wg.Wait()
		close(links)
	}()
	for l := range links {
		fmt.Println(l)
		go simpleCrawler(l.To, graph, links, wg)
	}

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
