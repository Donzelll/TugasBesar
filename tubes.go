package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	if t.Hour() < 12 {
		fmt.Println("Good morning!")
	} else if t.Hour() < 17 {
		fmt.Println("Good afternoon!")
	} else {
		fmt.Println("Good evening!")
	}
	fmt.Println("Bagas")
}
