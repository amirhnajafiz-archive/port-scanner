package main

import "fmt"

func main() {
	openports := Scan("")

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
