package rpl

// Receiver is used to receive Log on local.
// Acts as a Source.
type Receiver struct {
	targets []Target
	c       chan Log
}

func NewReceiver() Receiver {
	receiver := Receiver{}

	go func(r Receiver) {
		for {
			log := <-r.c
			for _, target := range r.targets {
				go func(t Target, l Log) {
					t.Writer() <- l
				}(target, log)
			}
		}
	}(receiver)

	return receiver
}

func (receiver Receiver) Register(target Target) {
	receiver.targets = append(receiver.targets, target)
}

func (receiver Receiver) Writer() chan<- Log {
	return receiver.c
}
