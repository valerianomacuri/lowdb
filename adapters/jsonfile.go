package adapters

type JSONFile[T interface{}] struct {
	adapter TextFile[T]
}

func NewJSONFile[T interface{}](filename string) *JSONFile[T] {
	return &JSONFile[T]{
		adapter: *NewTextFile[T](filename),
	}
}

func (j *JSONFile[T]) Read() T {
	return j.adapter.Read()
}

func (j *JSONFile[T]) Write(data T) {
	j.adapter.Write(data)
}
