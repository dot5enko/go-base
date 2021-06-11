package go_base

import (
	"go-base/errors"
	"runtime/debug"
)

type SafeExecutor struct {
	errors.ErrorNotifier
}

func (handler SafeExecutor) Handle(executor func() error) error {
	for {
		err := execWithRecovery(executor)
		if err == nil {
			return nil
		} else {
			handler.Notify(err)
		}
	}
}

func execWithRecovery(handler func() error) (err error) {

	defer func() {
		if x := recover(); x != nil {
			debug.PrintStack()
			err = errors.BasicError("Recovered event handling routine : %v", x)
		}
	}()

	handler()

	return
}
