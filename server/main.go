package main

import (
	handler "lite_queue_server/handler"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":6633")

	if err != nil {
		log.Fatalf("cannot listen: %v", err)
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Printf("cannot accept: %v", err)
			continue
		}

		h := handler.New(conn)

		go h.Handle()
	}
}
