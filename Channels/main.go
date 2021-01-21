package main

import "fmt"

func topla(s []int, c chan int) {
	toplam := 0
	for _, v := range s {
		toplam += v
	}
	c <- toplam //toplam, c kanalına yollandı
}
func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go topla(s[:len(s)/2], c)
	go topla(s[len(s)/2:], c)
	x, y := <-c, <-c //c kanalından veri alındı
	fmt.Println(x, y, x+y)
}
