package main

import (
	"sort"

	"github.com/amirhnajafiz/port-scanner/worker"
)

func Scan(uri string, portRange int) []int {
	ports := make(chan int, 100)
	response := make(chan int)

	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker.Worker(ports, response, uri)
	}

	go func() {
		for i := 0; i <= portRange; i++ {
			ports <- i
		}
	}()

	for i := 0; i <= portRange; i++ {
		port := <-response
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(response)

	sort.Ints(openports)

	return openports
}
