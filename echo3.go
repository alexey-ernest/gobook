package main

import (
	"strings"
	"fmt"
	"os"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}
