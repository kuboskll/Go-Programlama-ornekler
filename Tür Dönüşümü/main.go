package main

import "fmt"

func main() {
	a := 11.55
	b := int(a)
	fmt.Println(b)
}

/*
//Tür Dönüşümü
	var yazi = "11"
	var x = 5
	var f float64 = 2.0

	fmt.Println(yazi, x, f)
	//String olan yaziyi int yapalım
	sayi, _ := strconv.Atoi(yazi)
	fmt.Println("Sonuç(int): ", sayi)

	//İnt string ile yapma strconv.Itoa(sayi)

	fmt.Println("Sonuç(string): ", strconv.Itoa(sayi))


fmt.Println(sayi)
//toplam := sayi + 10
//fmt.Println(toplam)*/
