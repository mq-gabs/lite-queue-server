package lqs_client

import (
	"log"
	"lqs_client/connection"
)

func main() {
	c, err := connection.New()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer c.Close()

}
