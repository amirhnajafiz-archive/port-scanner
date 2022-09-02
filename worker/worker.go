package worker

import (
	"fmt"
	"log"
	"net"
)

// Worker makes a connection to the give uri and a port
// then it returns the result.
func Worker(port chan int, response chan int, uri string) {
	for p := range port {
		address := fmt.Sprintf("%s:%d", uri, p)

		conn, err := net.Dial("tcp", address)
		if err != nil {
			response <- 0

			continue
		}

		if err := conn.Close(); err != nil {
			log.Fatalf("error in connection close: %v\n", err)
		}

		response <- p
	}
}
