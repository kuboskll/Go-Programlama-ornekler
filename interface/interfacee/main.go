package main

import "fmt"

type inter interface {
	newinter()
}

type yapı struct {
	newyapı string
}

func (deneme yapı) newinter() {
	fmt.Println(deneme.newyapı)
}

func main() {
	var a inter = yapı{"Kübra"}
	a.newinter()
}
