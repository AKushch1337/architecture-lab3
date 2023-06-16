package painter

import (
	"image"
	"sync"

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
	Done     chan struct{}
	Stopped  bool
}

var size = image.Pt(800, 800)

// Start запускає цикл подій. Цей метод потрібно запустити до того, як викликати на ньому будь-які інші методи.
func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)
	l.MsgQueue = MessageQueue{}
	go l.processEvents()
}

func (l *Loop) processEvents() {
	for !l.Stopped || !l.MsgQueue.isEmpty() {
		op := l.MsgQueue.Pull()
		update := op.Do(l.next)
		if update {
			l.Receiver.Update(l.next)
			l.next, l.prev = l.prev, l.next
		}
	}
	close(l.Done)
}

// Post додає нову операцію у внутрішню чергу.
func (l *Loop) Post(op Operation) {
	if op != nil {
		l.MsgQueue.Push(op)
	}
}

// StopAndWait сигналізує
func (l *Loop) StopAndWait() {
	l.Post(OperationFunc(func(t screen.Texture) {
		l.Stopped = true
	}))
	<-l.Done
}

type MessageQueue struct {
	Queue   []Operation
	mu      sync.Mutex
	blocked chan struct{}
}

func (mq *MessageQueue) Push(op Operation) {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	mq.Queue = append(mq.Queue, op)

	if mq.blocked != nil {
		close(mq.blocked)
		mq.blocked = nil
	}
}

func (mq *MessageQueue) Pull() Operation {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if len(mq.Queue) == 0 {
		mq.blocked = make(chan struct{})
		mq.mu.Unlock()
		<-mq.blocked
		mq.mu.Lock()
	}
	op := mq.Queue[0]
	mq.Queue[0] = nil
	mq.Queue = mq.Queue[1:]
	return op
}
func (mq *MessageQueue) isEmpty() bool {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	return len(mq.Queue) == 0
}
