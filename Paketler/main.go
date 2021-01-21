package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {

	//fmt paketi
	fmt.Println("Hello")
	//rand paketi
	fmt.Println("Favori sayım", rand.Intn(10))
	//string arama varsa true döner
	fmt.Println(strings.Contains("Kübra", "br"))
	fmt.Println(strings.Count("Kübraaa", "a"))             //Kübra da kaç tane k var
	fmt.Println(strings.HasPrefix("Kübra_Kılıç", "Kübra")) //Ön ek
	fmt.Println(strings.HasSuffix("Kübra_kılıç", "kılıç")) //son ek arama
	fmt.Println(strings.Index("Kübra", "ü"))               // ü kaçıncı harf

}
