package main

import "fmt"

// func print() {
// 	var p *int
// 	fmt.Printf(*p)
// }
func main() {
	// 捕捉painc

	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	var i = 3
	var Slice [3]int

	fmt.Println(Slice[i])

}
