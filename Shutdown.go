package gobase

import (
	"os"
	"os/signal"
	"syscall"
)

// not safe, can raise panics
func OnShutdown(callback func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	callback()
}
