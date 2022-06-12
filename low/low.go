package low

type Adapter[T interface{}] interface {
	Read() T
	Write(data T)
}

type Low[T interface{}] struct {
	adapter Adapter[T]
	Data    T
}

func New[T interface{}](adapter Adapter[T]) *Low[T] {
	return &Low[T]{
		adapter: adapter,
	}
}

func (l *Low[T]) Read() {
	l.Data = l.adapter.Read()
}

func (l *Low[T]) Write() {
	l.adapter.Write(l.Data)
}
