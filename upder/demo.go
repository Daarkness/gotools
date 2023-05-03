package main

import "fmt"

var sz = 'a' - 'A'

func to_lower(x string) string {

	arr := []byte(x)

	for x := range arr {

		if 'A' <= arr[x] && arr[x] <= 'Z' {
			arr[x] = arr[x] + 'a' - 'A'
		}
	}
	return string(arr)
}

func to_upper(x string) string {
	arr := []byte(x)

	for x := range arr {

		if 'a' <= arr[x] && arr[x] <= 'z' {
			arr[x] = arr[x] - ('a' - 'A')
		}
	}
	return string(arr)

}

func main() {
	// var x int8
	// x = 127
	// x += 1
	// fmt.Println(x)
	// fmt.Println(unsafe.Sizeof(x))
	// y := "dsafasd\t234"
	// fmt.Println(y)

	s1 := "Hello World"
	fmt.Println(to_lower(s1))
	fmt.Println(to_upper(s1))
	// var c byte
	// c = s1[0]
	// s2 := s1[3:4]
	// fmt.Println(s1, c, s2)

	// s1[1] = 'c' 不能直接修改

}
