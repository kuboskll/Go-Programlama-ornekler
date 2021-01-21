package main

func main() {
	var i int
	for i = 0; i < 10; i++ {
		if i == 2 {
			continue
		} else if i == 4 {
			println("Sayı 4 oldu çıkış yapılıyor")
			break
		}
		println(i)
	}

}

/*//Swich Case Haftanın günleri
var gün int
fmt.Println("1-7 arası bir sayı giriniz: ")
fmt.Scan(&gün)
switch gün {
case 1:
	println("Pazartesi")
case 2:
	println("Salı")
case 3:
	println("Çarşamba")
case 4:
	println("Perşembe")
case 5:
	println("Cuma")
case 6:
	println("Cumartesi")
case 7:
	println("Pazar")
default:
	println("Yanlış değer girdiniz..")

}*/

//Klavyeden girilen iki sayıyı kontrol edelim
/*var a int
var b int
fmt.Println("a değerini giriniz")
fmt.Scan(&a)
fmt.Println("b değerini giriniz")
fmt.Scan(&b)
if a < b {
	println("b büyüktür")
} else if a > b {
	println("a büyüktür")
} else {
	println("İki sayı eşittir")
}

	   if c := 10; c == 1 {
		   println("c 1 dir")
	   } else {
		   println("c 10 dur")
	   }
   Dikkat edilmesi gereken C yi başka yerde kullanırsan hata alırsın
*/
