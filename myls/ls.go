package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../")
	if err != nil {
		log.Fatal(err)

	}
	infos, _ := f.Readdir(-1)
	// name, _ := f.Readdirnames(-1)
	for _, info := range infos {
		var flag int
		if info.IsDir() {
			flag = 1
		} else {
			flag = 0
		}
		fmt.Printf("%v %d %v %v\n", info.Mode(), flag, info.Size())
	}
}
