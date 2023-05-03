package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	// cmd := exec.Command("ls", "-l")
	// CombinedOutput  已经实现Wait
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(out))
	//========================================================

	//使用buf 分段读取器

	// cmd := exec.Command("ls", "-l")
	// out, _ := cmd.StdoutPipe()

	// f := bufio.NewReader(out)

	// if err := cmd.Start(); err != nil {
	// 	log.Fatal(err)
	// }
	// for {
	// 	line, err := f.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Print(line)
	// }

	// cmd.Wait()

	//========================================================

	//交互式调用

	// cmd := exec.Command("ls", "-l")

	// out, _ := cmd.StderrPipe()

	// buf := make([]byte, 1024)

	// for {
	// 	line, err := out.Read(buf)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Printf("read errr: %v", err)
	// 		break
	// 	}
	// 	fmt.Println(line)
	// }
	// time.Sleep(1 * time.Hour)
	// cmd := exec.Command("ls", "-l")
	// out, _ := cmd.StdoutPipe()

	// if err := cmd.Start(); err != nil {
	// 	log.Fatal(err)
	// }

	// f := bufio.NewReader(out)
	// for {
	// 	line, err := f.ReadString('\n')
	// 	if err != nil {
	// 		// log.Fatal(err)
	// 		break
	// 	}
	// 	fmt.Print(line)
	// }
	// time.Sleep(1 * time.Hour)
	// cmd.Wait() # 可以获取到退出信息

	//交互式调用

	// cmd := exec.Command("ls", "-l")

	// cmd := exec.Command("ls", "-l")
	// out, _ := cmd.StderrPipe()

	// buf := make([]byte, 1024)
	// fmt.Printf("start")

	// for {
	// 	_, err := out.Read(buf)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Printf("read errr: %v", err)
	// 		break
	// 	}

	// }
	// if err := cmd.Start(); err != nil {
	// 	log.Fatalf("start error :%s", err)
	// }
	// err := cmd.Wait()
	// if err != nil {
	// 	log.Fatalf("wait error :%s", err)
	// }

	//----------------------------------------------------

	f, err := os.Create("ls.out")

	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("ls", "-l")
	cmd.Stderr = f
	// cmd.Stdout = f
	fmt.Println("123")
	cmd.Run()

}
