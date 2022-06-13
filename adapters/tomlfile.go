package adapters

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml/v2"
)

type TOMLFile[T interface{}] struct {
	filename string
	writer   Writer
}

func NewTOMLFile[T interface{}](filename string) *TOMLFile[T] {
	return &TOMLFile[T]{
		filename: filename,
		writer:   *NewWriter(filename),
	}
}

func (j *TOMLFile[T]) Read() T {
	file, err := ioutil.ReadFile(j.filename)
	check(err)
	var data T
	err = toml.Unmarshal(file, &data)
	check(err)
	return data
}

func (j *TOMLFile[T]) Write(data T) {
	dataInBytes, err := toml.Marshal(data)
	check(err)
	j.writer.Write(dataInBytes)
}
