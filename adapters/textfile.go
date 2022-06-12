package adapters

import (
	"encoding/json"
	"io/ioutil"
)

type TextFile[T interface{}] struct {
	filename string
	writer   Writer[T]
}

func NewTextFile[T interface{}](filename string) *TextFile[T] {
	return &TextFile[T]{
		filename: filename,
		writer:   *NewWriter[T](filename),
	}
}

func (t *TextFile[T]) Read() T {
	file, err := ioutil.ReadFile(t.filename)
	check(err)
	var data T
	err = json.Unmarshal(file, &data)
	check(err)
	return data
}

func (t *TextFile[T]) Write(data T) {
	t.writer.Write(data)
}
