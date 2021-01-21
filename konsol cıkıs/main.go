package main

import "fmt"

func main() {

	fmt.Println("Adınızı Giriniz: ")
	var first string
	fmt.Scan(&first)
	fmt.Println("Soyadınızı giriniz: ")
	var second string
	fmt.Scan(&second)
	fmt.Print("Adınız ve Soyadınız: ")
	fmt.Print(first + " " + second)

}

/*fmt.Println("Bir sayi giriniz: ")
var sayi float64
fmt.Scan(&sayi)
fmt.Println("Bir sayi girinizzz: ")
var sayi1 float64
fmt.Scan(&sayi1)
var toplam float64 = (sayi + sayi1)
fmt.Println("İki sayının toplamı: ", toplam)*/

/*fmt.Println("Bir sayi giriniz: ")
var a int
fmt.Scan(&a)
fmt.Println("Bir sayi girinizzz: ")
var b int
fmt.Scan(&b)
var toplamm int = (a + b)
fmt.Println("İki sayının toplamı: ", toplamm)*/
