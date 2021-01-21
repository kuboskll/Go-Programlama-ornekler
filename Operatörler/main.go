package main

import "fmt"

//Operatörler
func main() {
	a := 10
	b := 30
	c := a + b
	//fmt.Println("c =", a+b)
	fmt.Println(c)
	c = c - 10
	fmt.Println(c)
	c *= 5
	fmt.Println(c)
	c = a % b
	fmt.Println(c)
	c++
	c--
	c++
	fmt.Println(c)

}
