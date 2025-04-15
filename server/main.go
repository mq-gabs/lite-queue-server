package main

import (
	handler "lite_queue_server/handler"
	"lite_queue_server/manager"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":6633")

	if err != nil {
		log.Fatalf("cannot listen: %v", err)
	}

	qm := manager.New()

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Printf("cannot accept: %v", err)
			continue
		}

		h := handler.New(conn, qm)

		go h.Handle()
	}
}
