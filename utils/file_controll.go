package utils

import (
	"io"
	"log"
	"os"
)

func CraeteDirectory(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func CreateFile(path string, filename string) *os.File {
	file, err := os.Create(path + filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func MultiWrite(out *os.File, file *os.File) io.Writer {
	multiOutput := io.MultiWriter(out, file)
	return multiOutput
}
