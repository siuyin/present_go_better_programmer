//010_OMIT
package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Hmm... You reported %q\nmain can help you.\n", r)
		}
	}()
	// Uncomment one below // HL
	//onlyPositiveMaster(-1)
	//onlyPositiveRookie(-1)
}

//020_OMIT
//030_OMIT
func onlyPositiveMaster(i int) {
	if i < 0 {
		log.Fatalf("Only positive numbers allowed: %v", i)
	}
}
func onlyPositiveRookie(i int) {
	if i < 0 {
		panic(fmt.Sprintf("Help! I got %v!", i))
	}
}

//040_OMIT
