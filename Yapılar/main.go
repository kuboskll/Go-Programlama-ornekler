package main

import (
	"fmt"
	"strconv"
)

type deneme struct {
	isim    string
	soyisim string
	yaş     int
}

//varsayılan ve boş yapıcı metot
func yenideneme() *deneme {
	c := new(deneme)
	return c
}

func main() {

	a := deneme{isim: "Kübra"}
	fmt.Println(a.isim)
	//2.yöntem
	b := new(deneme)
	b.isim = "Ayşe"
	fmt.Println(b.isim)
	//yapıcı metot kullanımı
	e := yenideneme()
	e.yaş = 22
	fmt.Println(e.yaş)
	//veriyi okuyalım
	d := new(deneme)
	d.isim = "Murat"
	d.soyisim = "Sönmez"
	d.yaş = 23
	var veri = d.isim + " " + d.soyisim + " " + strconv.Itoa(d.yaş) //yaşi dönüştürmek gerekiyor
	fmt.Println(veri)
}
