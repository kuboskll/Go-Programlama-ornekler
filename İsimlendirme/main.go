package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println(a, b, c)
	//Çıktıda otomatik ilk deger atamasını kendi 0 yapar

	var k, l, m int = 2, 3, 4
	fmt.Println(k, l, m)

	var mesaj string = "Hello Word! "
	var d, e, f = 1, true, 3.1
	fmt.Println(mesaj, d, e, f)

}

// Değişken tanımlama
/*
	var mesaj string
	mesaj = "Merhaba! "
	fmt.Println(mesaj)
*/

//Başka bir yolu
/*
	var mesaj string = "Selam go! "
	fmt.Println(mesaj)*/

//String tanımlamasak da olur
/*var mesaj = "Merhaba !"
fmt.Println(mesaj)
*/

// :=
/*a := 5
fmt.Println(a)*/
