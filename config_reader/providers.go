package config_reader

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

type (
	YamlProvider struct{}
	JsonProvider struct{}
)

func (p *YamlProvider) Unmarshal(source []byte, dest any) error {
	return yaml.Unmarshal(source, dest)
}

func (p *JsonProvider) Unmarshal(source []byte, dest any) error {
	return json.Unmarshal(source, dest)
}
