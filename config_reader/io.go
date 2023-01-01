package config_reader

import "os"

type IoHelper interface {
	ReadFile(string) ([]byte, error)
}

type IoReader struct{}

func (r *IoReader) ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
