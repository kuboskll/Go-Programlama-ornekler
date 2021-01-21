package main

import "fmt"

func main() {
	//Boş Tanımlayıcı _
	var _ int = 4
	b := 8
	fmt.Println(b)
	var _, a, d, c, _ int = 1, 2, 3, 4, 5
	fmt.Println("a=", a, "d=", d, "c=", c)

}
