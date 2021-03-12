package main

import "fmt"

func main() {
	longString := `
	{
		"propName": "value"
	}
	`
	fmt.Printf("string is %s", longString)
}
