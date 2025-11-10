package utils

import "fmt"

func Assert(cond bool, message string) {
	if !cond {
		fmt.Println(message)
	}
}
