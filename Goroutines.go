package gobase

import (
	"context"
	"time"
)

type ManagedRoutine struct {
	stepF func()
	endF  func()

	duration time.Duration

	ctx context.Context
}

func (r *ManagedRoutine) After(cb func()) *ManagedRoutine {
	r.endF = cb
	return r
}
func (r ManagedRoutine) Run() {
	end := false

	for !end {
		select {
		case <-time.After(r.duration):
			r.stepF()
		case <-r.ctx.Done():
			end = true
			if r.endF != nil {
				r.endF()
			}
		}
	}
}

func WithInterval(ctx context.Context, duration time.Duration, task func()) *ManagedRoutine {
	r := &ManagedRoutine{
		stepF:    task,
		endF:     nil,
		duration: duration,
		ctx:      ctx,
	}

	return r
}
