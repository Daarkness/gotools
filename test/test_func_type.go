package main

import (
	"errors"
	"fmt"
	"sort"
)

func add(m, n int) int {
	return m + n
}

func sub(m, n int) int {
	return m - n
}

//闭包
func addn(n int) func(int) int {
	return func(m int) int {
		return m + n
	}

}

//闭包生成器

func iter(s []int) func() (int, error) {
	var i = 0
	return func() (int, error) {
		if i >= len(s) {
			return 0, errors.New("end")
		}

		n := s[i]
		i += 1
		return n, nil

	}

}

func main() {
	// fmt.Println("test")
	// funcMap := map[string]func(int, int) int{
	// 	"add": add,
	// 	"sub": sub,
	// }

	// f := funcMap["add"]
	// fmt.Println(f(1, 3))

	// f := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Printf(">")
	// 	line, _ := f.ReadString('\n')
	// 	fmt.Print(line)
	// 	if line == "exit\n" {
	// 		break
	// 	}

	// }

	// s := strings.Map(func(r rune) rune {

	// 	return r - 32

	// }, "hello")
	// fmt.Println(s)

	// s1 := strings.Map(func(r rune) rune {
	// 	fmt.Println(r)
	// 	return r
	// }, "q231")
	// fmt.Println(s1)
	// f := addn(3)

	// fmt.Println(f(2))

	// x := addn(19)
	// fmt.Println(x(112))
	// f := (iter([]int{1, 2, 2}))
	// fmt.Println(f())

	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())

	//----------------------------------------
	//闭包的坑
	// var flist []func()

	// for i := 0; i < 3; i++ {
	// 	flist = append(flist, func() {
	// 		//
	// 		fmt.Println(i)
	// 	})
	// }

	// for _, f := range flist {
	// 	f()
	// }
	//----------------------------------------

	//排序

	type Student struct {
		Name string
		Id   int
	}

	s := []int{12, 3, 4, 1, 5}

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	ss := []Student{}
	ss = append(ss, Student{
		Name: "Aab",
		Id:   12,
	})
	ss = append(ss, Student{
		Name: "Jeck",
		Id:   1,
	})
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Id < ss[j].Id
	})
	fmt.Println(ss)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Name < ss[j].Name
	})
	fmt.Println(ss)
}
