package main

import "fmt"

func main() {
	//RANGE
	a := []string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println(i, "-", a[i])
	}
}

/*While Döngüsü
var i int = 1
for i < 10 {
	i += i
	println(i)
}
*/
//For Döngüsü
/*var i int
for i = 1; i < 10; i++ {
	fmt.Println("i değeri ", i)
}
*/
