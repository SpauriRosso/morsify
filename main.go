package main

import (
	"fmt"
	"log"
	"morsify/src/server"
	"runtime"
)

func main() {
	err := startServer(":2633")
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		txtErr := fmt.Sprintf("%v:%v %v", file, line, err)
		log.Fatalln(txtErr)
	}
}

func startServer(addr string) error {
	s := server.NewServer()
	return s.Start(addr)
}
