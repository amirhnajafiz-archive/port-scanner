package main

import "fmt"

func main() {
	openports := Scan("", 2048)

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
