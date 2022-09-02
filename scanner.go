package portscanner

import (
	"sort"

	"github.com/amirhnajafiz/port-scanner/worker"
)

const offset = 2048

// Scan takes an uri, portRange and number of workers and will
// test the uri ports in portRange (0, portRange) with workers and then
// returns the open ports.
func Scan(uri string, portRange []int, workers ...int) []int {
	numberOfWorkers := 100
	if len(workers) > 0 {
		numberOfWorkers = workers[0]
	}

	if len(portRange) == 0 {
		portRange = append(portRange, 0, offset)
	} else if len(portRange) == 1 {
		portRange = append(portRange, offset)
	}

	sort.Ints(portRange)

	ports := make(chan int, numberOfWorkers)
	response := make(chan int)

	defer func() {
		close(ports)
		close(response)
	}()

	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker.Worker(ports, response, uri)
	}

	go func() {
		for i := portRange[0]; i <= portRange[1]; i++ {
			ports <- i
		}
	}()

	for i := portRange[0]; i <= portRange[1]; i++ {
		port := <-response
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	sort.Ints(openPorts)

	return openPorts
}
