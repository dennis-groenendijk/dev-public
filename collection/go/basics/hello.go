package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Dennis", "name of the person")

	var number int
	flag.IntVar(&number, "number", 1, "number of repetition")

	flag.Parse()

	for i := 0; i < number; i++ {
		fmt.Println("Hello", name, "!")
	}
}
