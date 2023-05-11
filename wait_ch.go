package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func waitForSignal(stopCh chan struct{}) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh
	close(stopCh)
}

func main() {
	stopCh := make(chan struct{})
	go waitForSignal(stopCh)

	// Do something else until the signal is received
	// time.Sleep(5 * time.Second)
	fmt.Println("Wait")

	<-stopCh
}
