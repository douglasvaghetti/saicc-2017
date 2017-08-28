package main

import "testing"

func Test_globalize(t *testing.T) {
	type args struct {
		baseDomain string
		from       string
		to         string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"basico", args{"http://devopers.com.br", "http://devopers.com.br", "http://devopers.com.br/home/"}, "http://devopers.com.br/home/"},
		{"caminho relativo", args{"http://devopers.com.br", "http://devopers.com.br", "/home/"}, "http://devopers.com.br/home/"},
		{"outro dominio", args{"http://pudim.com.br", "http://pudim.com.br/asd", "http://devopers.com.br/"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := globalize(tt.args.baseDomain, tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("globalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
