package adapters

import (
	"io/ioutil"
	"os"

	"path"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Writer struct {
	filename     string
	tempFilename string
}

func NewWriter(filename string) *Writer {
	return &Writer{
		filename:     filename,
		tempFilename: getTempFilename(filename),
	}
}

func getTempFilename(file string) string {
	return path.Join(path.Dir(file), "."+path.Base(file)+".tmp")
}
func (w *Writer) Write(data []byte) {
	err := ioutil.WriteFile(w.tempFilename, data, 0644)
	check(err)
	err = os.Rename(w.tempFilename, w.filename)
	check(err)
}
