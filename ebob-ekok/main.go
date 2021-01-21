package main

import "fmt"

func main() {
	var sayi1 int = 48
	var sayi2 int = 160
	var ebob int = 1
	var ekok int
	fmt.Println("Birinci sayıyı giriniz: ")
	fmt.Scan(&sayi1)
	fmt.Println("İkinci sayıyı giriniz: ")
	fmt.Scan(&sayi2)

	for i := 1; i <= sayi1 && i <= sayi2; i++ {
		if sayi1%i == 0 && sayi2%i == 0 {
			ebob = i
		}

		ekok = (sayi1 * sayi2) / ebob

	}
	fmt.Println("Ekok:", ekok)
	fmt.Println("Ebob: ", ebob)
}
