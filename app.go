package main

import "fmt"

func Run() {
	fmt.Println("vim-go")
}

//go:generate counterfeiter -o fakes/fake_x.go . X
type X interface {
	String() error
}
