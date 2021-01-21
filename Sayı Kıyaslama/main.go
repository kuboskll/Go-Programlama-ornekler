package main

import "fmt"

func main() {
	var sayi1, sayi2 int
	fmt.Println("1.Sayıyı giriniz: ")
	fmt.Scan(&sayi1)
	fmt.Println("2. Sayıyı giriniz: ")
	fmt.Scan(&sayi2)
	if sayi1 < sayi2 {
		fmt.Println("1.Sayı küçüktür...")
	} else if sayi1 > sayi2 {
		fmt.Println("2. Sayı küçüktür... ")
	} else {
		fmt.Println("İki sayı eşittir...")
	}
}
