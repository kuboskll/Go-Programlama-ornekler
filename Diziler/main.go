package main

import "fmt"

func main() {
	myMap := make(map[int]string)
	myMap[06] = "Ankara"
	myMap[34] = "İstanbul"
	myMap[81] = "Düzce"
	myMap[41] = "Kocaeli"
	myMap[16] = "Bursa"
	myMap[01] = "Adana"
	fmt.Println(myMap[01])
	fmt.Println(myMap[34])
	delete(myMap, 06)
	fmt.Println("Ankara Silindi")
	fmt.Println(myMap)
	for k, v := range myMap {
		fmt.Println(k, v)
	}
}

/* myArray := [6]string{"Ali", "Ayşe", "Murat", "Fatih", "Emre", "Buse"}
Slice1 := myArray[1:5]
fmt.Println(myArray)
fmt.Println(Slice1)
fmt.Println(len(Slice1))
fmt.Println(cap(Slice1))
var myArray2 []int
Slice2 := myArray2[1:2]
fmt.Println(Slice2)



var deneme []string
	if deneme == nil {
		fmt.Println("Boş Dilim!")
	}*/
