package main

import (
	"fmt"
	"time"
)

//010_OMIT
func a() {
	go func() {
		for {
			fmt.Print("A")
			time.Sleep(1 * time.Second) // this is blocking!
		}
	}()
}
func b() {
	go func() {
		for {
			fmt.Print("b")
			time.Sleep(1 * time.Second) // this is blocking!
		}
	}()
}

func main() {
	a()
	b()
	select {} // wait forever
}

//020_OMIT
