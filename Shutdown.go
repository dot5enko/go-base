package gobase

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// not safe, can raise panics
func OnShutdown(callback func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	callback()
}

func ExecuteWithInterval(ctx context.Context, duration time.Duration, task func()) {
	end := false

	for !end {
		select {
		case <-time.After(duration):
			task()
		case <-ctx.Done():
			end = true
		}
	}
}
