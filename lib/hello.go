package lib

import (
	"fmt"
)

func Hello(target string) {
	if target == "toy" {
		fmt.Printf("That's kind to greet me! :)\n")
	} else {
		fmt.Printf("Hello %v!\n", target)
	}
}
