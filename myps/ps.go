package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func isNum(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true

}

func main() {
	f, err := os.Open("/proc")
	if err != nil {
		log.Fatal(err)

	}
	infos, _ := f.Readdir(-1)
	// name, _ := f.Readdirnames(-1)
	fmt.Printf("%s %10s\n", "USER", "PID")
	for _, info := range infos {

		if isNum(info.Name()) {
			fmt.Println(info.Name())
			// fmt.Print(info.Sys())
			// for x, v := range info.Sys() {
			// 	fmt.Println(x, v)
			// }
			x := info.Sys()
			fmt.Printf("info.Sys(): %v\n", x)

		}

	}

}
