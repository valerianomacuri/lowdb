package adapters

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type YAMLFile[T interface{}] struct {
	filename string
	writer   Writer
}

func NewYAMLFile[T interface{}](filename string) *YAMLFile[T] {
	return &YAMLFile[T]{
		filename: filename,
		writer:   *NewWriter(filename),
	}
}

func (j *YAMLFile[T]) Read() T {
	file, err := ioutil.ReadFile(j.filename)
	check(err)
	var data T
	err = yaml.Unmarshal(file, &data)
	check(err)
	return data
}

func (j *YAMLFile[T]) Write(data T) {
	dataInBytes, err := yaml.Marshal(data)
	check(err)
	j.writer.Write(dataInBytes)
}
