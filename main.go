package main

import (
	"fmt"

	"github.com/9bany/kgs/pkg/key0"
)

func main() {
	gen := key0.NewGen(key0.WithLen(9))
	k, _ := gen.New()
	fmt.Println(k)
}
