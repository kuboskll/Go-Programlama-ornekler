package main

import "fmt"

func main() {
	var sayi int
	fmt.Println("Lütfen sayıyı giriniz: ")
	fmt.Scan(&sayi)
	if sayi == 0 {
		fmt.Println("Sayınız sıfırdır...")
	} else if sayi%2 == 0 {
		fmt.Println("Sayınız Çifttir...")
	} else {
		fmt.Println("Sayınız Tektir...")
	}
}
