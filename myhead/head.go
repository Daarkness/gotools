package main

import (
	"fmt"
)

// var (
// 	line     = flag.Int("n", 10, "num")
// 	filename = flag.String("f", "a.txt", "filename")
// )

func Write99() {
	x := 10
	for i := range x {
		fmt.Printf(i)
	}

}

func main() {
	Write99()

	// flag.Parse()
	// fmt.Print(*line)
	// fmt.Print(*filename)
}
