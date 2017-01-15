package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//010_OMIT
func main() {
	rf, _ := os.Open("rdr/mydat.txt")
	rs := strings.NewReader("Hello World!\n")
	read(rf, "file")
	read(rs, "string")
}
func read(r io.Reader, src string) {
	p := make([]byte, 100)
	n, _ := r.Read(p)
	fmt.Printf("%d bytes read from %s\n", n, src)
	fmt.Printf("%s\n", p)
}

//020_OMIT
