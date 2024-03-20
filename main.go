package main

import (
	"fmt"

	keys "github.com/9bany/kgs/pkg/keys"
)

func main() {
	gen := keys.NewAlphabet()

	ch := make(chan string, 10)
	done := make(chan bool)

	go gen.Gen(3, ch, done)

	for {
		select {
		case key := <-ch:
			fmt.Println("sent message", key)
		case <-done:
			fmt.Println("end")
			return
		}
	}
}
