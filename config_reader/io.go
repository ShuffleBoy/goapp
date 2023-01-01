package config_reader

type IoHelper interface {
	ReadFile(string) ([]byte, error)
}
