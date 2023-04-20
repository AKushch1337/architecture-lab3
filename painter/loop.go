package painter

import (
	"image"

	"golang.org/x/exp/shiny/screen"
)

// Receiver отримує текстуру, яка була підготовлена в результаті виконання команд у циелі подій.
type Receiver interface {
	Update(t screen.Texture)
}

// Loop реалізує цикл подій для формування текстури отриманої через виконання операцій отриманих з внутрішньої черги.
type Loop struct {
	Receiver Receiver

	next screen.Texture // текстура, яка зараз формується
	prev screen.Texture // текстура, яка була відправленя останнього разу у Receiver

	MsgQueue MessageQueue
}

var size = image.Pt(400, 400)

// Start запускає цикл подій. Цей метод потрібно запустити до того, як викликати на ньому будь-які інші методи.
func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)
	l.MsgQueue = MessageQueue{}
	go l.processEvents()
}

func (l *Loop) processEvents() {
	for {
		if op := l.MsgQueue.Pull(); op != nil {
			if update := op.Do(l.next); update {
				l.Receiver.Update(l.next)
				l.next, l.prev = l.prev, l.next
			}
		}
	}
}

// Post додає нову операцію у внутрішню чергу.
func (l *Loop) Post(op Operation) {
	if op != nil {
		l.MsgQueue.Push(op)
	}
}

// StopAndWait сигналізує
func (l *Loop) StopAndWait() {

}

type MessageQueue struct {
	Queue []Operation
}

func (mq *MessageQueue) Push(op Operation) {
	mq.Queue = append(mq.Queue, op)
}

func (mq *MessageQueue) Pull() Operation {
	if len(mq.Queue) == 0 {
		return nil
	}
	op := mq.Queue[0]
	mq.Queue = mq.Queue[1:]
	return op
}
