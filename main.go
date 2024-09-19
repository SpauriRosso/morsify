package main

import (
	"fmt"
	"github.com/SpauriRosso/dotlog"
	"morsify/src/server"
	"runtime"
)

var addr string

func main() {
	addr = ":2633"
	launchTxt := fmt.Sprintf("Server started at %v", addr)
	dotlog.Info(launchTxt)
	err := startServer(addr)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		txtErr := fmt.Sprintf("%v:%v %v", file, line, err)
		dotlog.Error(txtErr)
	}
}

func startServer(addr string) error {
	s := server.NewServer()
	return s.Start(addr)
}
