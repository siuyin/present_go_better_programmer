package main

import (
	"fmt"
	"strings"
)

func main() {
	p := make([]byte, 100)
	strRdr := strings.NewReader("Hello World!\n")
	n, _ := strRdr.Read(p)
	fmt.Printf("%d bytes read\n", n)
	fmt.Printf("%s\n", p)
}
