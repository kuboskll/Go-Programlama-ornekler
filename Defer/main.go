package main

import "fmt"

func main() {
	//DEFER
	defer fmt.Println("Normal şartlarda ilk çıktı")
	fmt.Println("Fakat ilk derlenen benim")
	defer fmt.Println("Merhaba")
	fmt.Println("Nomal şartlarda en son derlenen")
}
