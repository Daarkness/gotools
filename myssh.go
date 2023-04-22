package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/kr/pty"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	addr    = flag.String("addr", ":8080", "address.")
	iserver = flag.Bool("s", false, " run as sever.")
)

func handle(conn net.Conn) {
	defer conn.Close()
	log.Printf("[login:] %s", conn.RemoteAddr())

	cmd := exec.Command("sh")
	tty, err := pty.Start(cmd)
	if err != nil {
		log.Print(err)
		return
	}
	go io.Copy(tty, conn)
	go io.Copy(conn, tty)
	cmd.Wait()
	log.Printf("[logout]: %s", conn.RemoteAddr())

}

func server() {
	l, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handle(conn)

	}
}

func client() {
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer terminal.Restore(0, oldState)
	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)

}

func main() {
	flag.Parse()
	if *iserver {
		server()
	} else {
		client()
	}
}
