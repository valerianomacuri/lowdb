package low

// Adapter for save in different formats
type Adapter[T interface{}] interface {
	Read() T
	Write(data T)
}

// Low receives the adapter (json, etc...) and the generic data type that will be used to parse the database
type Low[T interface{}] struct {
	adapter Adapter[T]
	Data    T
}

// New create a new instance of Low
func New[T interface{}](adapter Adapter[T]) *Low[T] {
	return &Low[T]{
		adapter: adapter,
	}
}

// Read connect to json database
func (l *Low[T]) Read() {
	l.Data = l.adapter.Read()
}

// Write into json database
func (l *Low[T]) Write() {
	l.adapter.Write(l.Data)
}
