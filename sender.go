package rpl

// Sender acts as a Target
// to send Log from Source to local.
// Works on remote.
type Sender struct {
	c chan Log
}

func NewSender() Sender {
	return Sender{}
}

func (sender Sender) Writer() chan<- Log {
	return sender.c
}

func (sender Sender) Reader() <-chan Log {
	return sender.c
}
