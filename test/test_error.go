package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// func read(f *os.File) (string, error) {
// 	buf, err := ioutil.ReadAll(f)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(buf), nil
// }

func read(f *os.File) (string, error) {
	var total []byte
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		total = append(total, buf[:n]...)
	}
	return string(total), nil
	// defer f.Close()

}

func main() {

	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	str, err := read(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
	//错误的类型
	//1 出现就推出 log.Fatal()
	//2 需要处理的err  直接处理
	//3 向上抛出的err  return XXX err

}
