package errors

type ErrorNotifier struct {
	notifier func(err error)
}

func (receiver *ErrorNotifier) SetNotifier(ehandler func(err error)) {
	receiver.notifier = ehandler
}
func (receiver ErrorNotifier) Notify(err error) {
	if receiver.notifier != nil {
		go receiver.notifier(err)
	}
}
