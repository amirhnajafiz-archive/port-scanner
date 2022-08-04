package worker

import (
	"fmt"
	"net"
)

type Worker struct {
	URI string
}

func (w *Worker) Start(port chan int, response chan int) {
	for p := range port {
		address := fmt.Sprintf("%s:%d", w.URI, p)

		conn, err := net.Dial("tcp", address)
		if err != nil {
			response <- 0
			continue
		}

		conn.Close()
		response <- p
	}
}
