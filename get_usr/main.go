package main

import "fmt"

//010_OMIT

// Getter returns a string associated with record number i.
type Getter interface {
	Get(i int) string
}
type myFakeUserDB struct{}

func (fdb myFakeUserDB) Get(i int) string {
	return fmt.Sprintf("User %d is 'SiuYin'", i)
}

func main() {
	g := myFakeUserDB{} // g can be any type that has a Get(i int) string method
	fmt.Println(g.Get(1))
}

//020_OMIT
