package rpl

// Receiver is used to receive Log on local.
// Acts as a Source.
type Receiver struct {
	targets []*Target
	c       chan Log

	// ChOffset is the offset of [Log.Ch].
	ChOffset uint16
}

func NewReceiver() *Receiver {
	receiver := Receiver{}

	go func(r *Receiver) {
		for {
			originLog := <-r.c
			log := Log{
				Ch:    originLog.Ch + r.ChOffset,
				Level: originLog.Level,
				Value: originLog.Value,
			}
			for _, target := range r.targets {
				go func(t *Target, l Log) {
					(*t).Writer() <- l
				}(target, log)
			}
		}
	}(&receiver)

	return &receiver
}

func (receiver *Receiver) Register(target *Target) {
	receiver.targets = append(receiver.targets, target)
}

func (receiver *Receiver) Writer() chan<- Log {
	return receiver.c
}
