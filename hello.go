package lib

import (
	"fmt"
)

func hello(i int) {
	if i > 0 {
		fmt.Println("Hello world!")
	} else {
		fmt.Println("Hello!")
	}
}
