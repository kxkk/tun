package main

import (
	"log"

	"github.com/4396/tun/proxy/tcp"
	"github.com/4396/tun/server"
)

func main() {
	p, err := tcp.Listen(":7070")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.ListenAndServe(":8867", p))
}
