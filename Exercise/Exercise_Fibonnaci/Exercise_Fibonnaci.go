package main

import (
	"fmt"
)

func main() {
	f := fibonnaci()
	for i:= 0; i < 10; i++ {
		fmt.Println(f())
	}

}

func fibonnaci() func() int {
	 a,b := 0,1
	return func() int {
		if a <= 0 {
			a,b = b, a+b
			return 0
		}
		a, b = b, a+b
		return a
	}
}