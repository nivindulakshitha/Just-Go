package main

import "fmt"

func Print(text string, method string) {
	switch method {
	case "same-line":
		fmt.Print(text)
		break
	default:
		fmt.Println(text)
	}
}