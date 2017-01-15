package main

import (
	"fmt"
	"time"
)

//010_OMIT
func a() <-chan bool {
	ch := make(chan bool)
	go func() {
		for {
			fmt.Print("A")
			ch <- true
			time.Sleep(1 * time.Second)
		}
	}()
	return ch
}

//020_OMIT
//030_OMIT
func b(ch <-chan bool) {
	go func() {
		for {
			select { // not strictly needed, but here to illustrate use of select
			case <-ch:
				fmt.Print("b")
			}
		}
	}()
}

//040_OMIT
//050_OMIT
func main() {
	ch := a() // recall -- a() returns a chan of bool
	b(ch)     // we connect this channel to b
	select {}
}

//060_OMIT
