package adapters

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"path"
)

type Writer[T interface{}] struct {
	filename     string
	tempFilename string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NewWriter[T interface{}](filename string) *Writer[T] {
	return &Writer[T]{
		filename:     filename,
		tempFilename: getTempFilename(filename),
	}
}

func getTempFilename(file string) string {
	return path.Join(path.Dir(file), "."+path.Base(file)+".tmp")
}
func (w *Writer[T]) Write(data T) {
	file, _ := json.MarshalIndent(data, "", "	")
	err := ioutil.WriteFile(w.tempFilename, file, 0644)
	check(err)
	err = os.Rename(w.tempFilename, w.filename)
	check(err)
}
