package main

import (
	"fmt"
	"os"
)

func main() {
	p := make([]byte, 100) // make byte slice with capacity of 100 bytes.
	file, _ := os.Open("rdr/mydat.txt")
	n, _ := file.Read(p)
	fmt.Printf("%d bytes read\n", n)
	fmt.Printf("%s\n", p)
}
