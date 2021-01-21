package main

import (
	"fmt"
	"time"
)

func main() {
	tarihdeneme := time.Now()
	fmt.Println(tarihdeneme)
	t := time.Date(2018, time.December, 10, 2, 0, 0, 0, time.UTC)
	fmt.Println(t)
	fmt.Println("MONTH: ", t.Month())
	fmt.Println("DAY: ", t.Day())
	fmt.Println("WEEKDAY: ", t.Weekday())
}
