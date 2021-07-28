package event

import (
	"log"

	"github.com/google/uuid"
)

type CbFunc func(arg ...interface{})

var (
	callbacks map[string]map[string]CbFunc = make(map[string]map[string]CbFunc)
)

func Emit(name string, arg ...interface{}) {
	for _, it := range callbacks[name] {
		go func(cbref CbFunc) {

			defer func() {
				recovered := recover()
				if recovered != nil {
					log.Printf("Unable to execute cb callback: %v", recovered)
				}
			}()

			cbref(arg...)

		}(it)
	}
}

func On(name string, cb CbFunc) string {

	uid := uuid.NewString()

	_, ok := callbacks[name]
	if !ok {
		callbacks[name] = make(map[string]CbFunc)
	}

	callbacks[name][uid] = cb

	return uid
}

func Cancel(id string) bool {
	panic("cancel not implemented")
	return false
}
