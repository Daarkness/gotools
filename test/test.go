// package main

// func WriteLine(filename string, line string) {
// 	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	f.WriteString(line)
// 	defer f.Close()
// }

// func main() {
// 	var x, y int
// 	for x = 1; x < 10; x++ {
// 		// fmt.Println(x)
// 		var line string
// 		for y = 1; y <= x; y++ {
// 			fmt.Printf("%v *%v \t", x, y)
// 			line += fmt.Sprintf("%v * %v \t", x, y)
// 		}
// 		line += "\n"
// 		// fmt.Print(line)
// 		WriteLine("9x9.txt", line)
// 		line = ""
// 		fmt.Println("")
// 	}

// }

// func main() {
// var a [3]int
// fmt.Println(a[0])
// fmt.Println(a[len(a)-1])
// for i, v := range a {

// 	fmt.Printf("i %d, v %d\n", i, v)
// }

// for _, v := range a {

// 	fmt.Printf("v %d\n", v)
// }
//初始化

// var b [3]int = [3]int{1, 2, 3}

// var c [3]int = [3]int{1, 2} //没有赋值的 默认为0
// fmt.Println(b, c)

// q1 := [...]int{1, 2, 3, 4} //不指定长度
// fmt.Println(q1)
// q2 := [...]int{4: 2, 10: -1} //指定下标复制，没有指定的为0
// fmt.Println(q2)
//复制

// a1 := [3]int{1, 2, 3}
// fmt.Print(a1)

// var a2 [3]int

// a2 = a1
// fmt.Printf("a1:%v , a2: %v", &a1[0], &a1[2])
// fmt.Println(unsafe.Sizeof(a1))
// var x, y *int
// y = &a1[0]
// x = y //对于指针而言，他们的值就是地址
// fmt.Println(x, y)
// fmt.Println(&x, &y)

//md5
// data := []byte("hello")

// md5sum := md5.Sum(data)

// fmt.Printf("%x\n", md5sum)

// var b []byte
// b = []byte("this is test")

// fmt.Println(b)
// fmt.Println(string(b))

//切片
// primes := [6]int{2, 3, 5, 6, 7, 8}

// var s []int = primes[1:4] //s[0] 指向了  primes[1]
// fmt.Println(s, primes)
// fmt.Println(&s[0], &primes[1])
// s[0] = 12
// fmt.Println(s, primes])

// name := [4]string{
// 	"Jonh",
// 	"paul",
// 	"George",
// 	"Ringo",
// }

// fmt.Println(name)

// a := name[0:2]
// b := name[1:3]
// fmt.Println(name)
// fmt.Println(a, b)
// b[0] = "XXX"
// fmt.Println(a, b)
// fmt.Println(name)

// fmt.Println(a)

//切片声名
// s := []int{1, 3, 4}
// s = append(s, 122)
// fmt.Println(s, len(s), cap(s))

// fmt.Println(s[:])
// fmt.Println(s[1:])
// fmt.Println(s[:3])
// fmt.Println(s[:0])

// fmt.Println(s[1:])
// fmt.Println(s, len(s), cap(s))
// s = s[:6]
// fmt.Println(s, len(s), cap(s))
// s = append(s, 32)
// fmt.Println(s, len(s), cap(s))
// fmt.Println(s[:0])

//空指针切片

// f := []int{123}

// s := f[:0]
// fmt.Println(s, len(s), cap(s))
// if s == nil {
// 	fmt.Println("nil")
// }

//make 创建切片

// a := make([]int, 5)
// fmt.Println("a:", a)

// b := make([]int, 1, 5) //长度和初始化长度
// fmt.Println("b", b)
// c := b[:5]
// fmt.Println("c", c)

//切片扩容
// a := make([]int, 0)
// a = append(a, 3)

// fmt.Println("a", a)
// a = append(a, 3, 3, 3, 4)

// fmt.Println("a", a)
// _ = append(a, 232)

// s := make([]int, 5)

// s1 := []int{2, 4, 5}

// s = append(s, s1...)

// fmt.Println(s)

// a := make([]int, 5)
// PrintSlice("a", a)
// b := make([]int, 0, 5)
// PrintSlice("b", b)
// c := b[:2]
// PrintSlice("c", c)
// d := c[2:5]
// PrintSlice("d", d)
//11 7 5 3 2
// 3
// 5 7 11 2 3

// args := make(map[string]int)
// args["a"] = 1
// args["b"] = 3

// for k, v := range args {
// 	fmt.Println(k, v)
// }

// args["a"] = args["b"] + 2
// for k, v := range args {
// 	fmt.Println(k, v)
// }
// c, ok := args["c"]
// if ok {
// 	fmt.Print(c)
// } else {
// 	fmt.Println("not found.")
// }

// if c, ok := args["c"]; ok {
// 	fmt.Println(c)
// }
// delete(args, "a")
// for k, v := range args {
// 	fmt.Println(k, v)
// }
// fmt.Println(strings.Fields("aa bb  cc"))

// }

// func PrintSlice(s string, x []int) {
// 	fmt.Printf("%s len=> %d cap=> %d %v\n", s, len(x), cap(x), x)

// }

// type Student struct {
// 	Id   int
// 	Name string
// }

// func main() {
// 	var s Student
// 	s.Id = 1
// 	s.Name = "Jack"
// 	fmt.Println(s)

// 	s1 := Student{
// 		Id:   2,
// 		Name: "Alice",
// 	}
// 	fmt.Println(s1)

// 	s1 = s
// 	fmt.Println(s1)

// }
