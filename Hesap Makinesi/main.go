package main

import "fmt"

func main() {
	fmt.Println("Toplama için 1 basınız")
	fmt.Println("Çıkarma için 2 basınız")
	fmt.Println("Çarpma için 3 basınız")
	fmt.Println("Bölme için 4 basınız")
	var sayi1, sayi2 int
	var secim int
	fmt.Println("İki sayı giriniz: ")
	fmt.Scan(&sayi1)
	fmt.Scan(&sayi2)
	fmt.Println("Seçim yapınız: ")
	fmt.Scan(&secim)
	switch secim {
	case 1:
		fmt.Println("Toplamanın sonucu: ", sayi1+sayi2)
	case 2:
		fmt.Println("Çıkarmanın sonucu: ", sayi1-sayi2)
	case 3:
		fmt.Println("Çarpmanın sonucu: ", sayi1*sayi2)
	case 4:
		fmt.Println("Bölmenin sonucu: ", sayi1/sayi2)
	default:
		fmt.Println("Yanlış secim yaptınız...")
	}
}
