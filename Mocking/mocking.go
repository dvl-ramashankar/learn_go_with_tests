package main

import (
	"fmt"
	"time"
)

const finalWord = "Go!"

func main() {
	countdown(3)
}

func countdown(x int) {
	for i := x; i >= 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	fmt.Println(finalWord)
}
