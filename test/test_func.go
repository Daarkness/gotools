package main

import "fmt"

//多返回值
func Swap(x, y string) (string, string) {
	// 返回值的声明必须用括号且用,隔开
	return y, x

}

//命令返回值
func split(sum int) (x, y int) {

	x = sum / 10
	y = sum % 10
	return

}

// 可变参数

func sum(args ...int) int {
	n := 0
	for i := 0; i < len(args); i++ {
		n += args[i]
	}
	return n
}

//递归

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {

	a, b := Swap("s234", "213")
	fmt.Println(a, b)

	fmt.Println(split(101))
	fmt.Println(sum(32, 33, 44))
	fmt.Println("-------")
	fmt.Println(fib(10))
}
