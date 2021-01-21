package main

import "fmt"

func main() {
	var faktoriyel int = 1
	var sayi1 int
	fmt.Println("Sayı giriniz: ")
	fmt.Scan(&sayi1)
	for i := 2; i <= sayi1; i++ {
		faktoriyel *= i
	}
	fmt.Println("Girilen sayının faktoriyeli: ", faktoriyel)
}
