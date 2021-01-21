package main

import "fmt"

func main() {
	i := 0
	var sayi int
	var sayac int = 0
	fmt.Println("Bir sayi giriniz: ")
	fmt.Scan(&sayi)
	for i = 2; i < sayi; i++ {
		if sayi%i == 0 {
			sayac++
			break
		}
	}
	if sayac == 0 {
		fmt.Println("Asal sayıdır")
	} else {
		fmt.Println("Asal sayı değildir")
	}
}
