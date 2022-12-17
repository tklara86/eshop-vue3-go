package main

import (
	"log"

	"github.com/tklara86/eshop-vue3-go/internal/server"
)

func main() {

	s, err := server.New(server.Config{
		Port: 9090,
	})
	if err != nil {
		log.Fatalf("server not created. Exited with error: %s", err)
	}
	err = s.Run()
	if err != nil {
		log.Fatalf("sever did not start: %s", err)
	}

}
