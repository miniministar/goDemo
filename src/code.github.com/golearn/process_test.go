package main

import (
	"fmt"
	"testing"
)

func TestProcess(t *testing.T) {
	var a int = 9
	if a > 9 {
		fmt.Println("a>9")
	} else if a == 9 {
		fmt.Println("a==9")
	} else {
		fmt.Println("a<9")
	}

	for i := 10; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("1" + "2")
	for i := 1; i <= 9; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("%d\n", i)
			fmt.Printf("%d*%d=%d", i, j, i*j)
			if j < i-1 {
				fmt.Printf(",")
			} else if j == i-1 {
				fmt.Println()
			}
		}
	}
}
