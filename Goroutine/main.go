package main

import (
	"time"
)

func alfabe() {
	for x := byte('a'); x <= byte('z'); x++ {
		println(string(x))
	}
}
func main() {
	go alfabe()
	time.Sleep(100 * time.Millisecond)
	//go yazmadan arasındaki fark işlemi tek tek satır satır yapar
	// go ekiyle ise yeni bir tread oluşturcak ve ana program durmadan devam eder.
}
