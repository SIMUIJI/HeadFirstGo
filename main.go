package main

import (
	"fmt"
	"work_go/src/github.com/headfirstgo/magazine"
)

func main() {
	var s magazine.Subscriber
	s.Rate = 4.99

	fmt.Println(s.Rate)
}
