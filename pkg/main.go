package main

import (
	"fmt"

	"./cfg"
	"github.com/siuyin/dflt"
)

func main() {
	port := dflt.EnvString("PORT", "8080")
	fmt.Println(port, cfg.MyName) // if PORT env var is unset, prints "8080 SiuYin"
}
