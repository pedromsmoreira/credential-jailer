package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	reader := &EnvVarReader{
		proxy: &OsProxy{},
	}

	server := NewServer(reader)
	err := server.Start("4000")

	if err != nil {
		fmt.Printf("Error starting the server. Error: %v", err)
		os.Exit(1)
	}
	ss := make(chan os.Signal, 1)
	fmt.Println("Sidelogger is running. Ctrl-C to exit")
	signal.Notify(ss, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-ss
	fmt.Println("Sidelogger is terminating!")
	err = server.Shutdown()
	if err != nil {
		fmt.Printf("Failed to shutdown API server: %v", err)
	}
	fmt.Println("Sidelogger terminated!")

	os.Exit(0)
}
