package main

import (
	"fmt"
	"os"
)

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[rico@%s]$", host)

}
