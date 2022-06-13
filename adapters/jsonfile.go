package adapters

import (
	"encoding/json"
	"io/ioutil"
)

type JSONFile[T interface{}] struct {
	filename string
	writer   Writer
}

func NewJSONFile[T interface{}](filename string) *JSONFile[T] {
	return &JSONFile[T]{
		filename: filename,
		writer:   *NewWriter(filename),
	}
}

func (j *JSONFile[T]) Read() T {
	file, err := ioutil.ReadFile(j.filename)
	check(err)
	var data T
	err = json.Unmarshal(file, &data)
	check(err)
	return data
}

func (j *JSONFile[T]) Write(data T) {
	dataInBytes, err := json.MarshalIndent(data, "", "	")
	check(err)
	j.writer.Write(dataInBytes)
}
