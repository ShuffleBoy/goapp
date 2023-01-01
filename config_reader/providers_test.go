package config_reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	yamlExample1 = `
test:
  omg: 1
`
	jsonExample1 = `{"test":{"omg": 1}}`
)

type (
	testStruct struct {
		Test struct {
			Omg int `yaml:"omg"`
		} `yaml:"test"`
	}
)

func TestYamlProvider_Unmarshal(t *testing.T) {
	p := new(YamlProvider)
	dest := new(testStruct)
	err := p.Unmarshal([]byte(yamlExample1), dest)
	assert.NoError(t, err)
	assert.Equal(t, 1, dest.Test.Omg)
}

func TestJsonProvider_Unmarshal(t *testing.T) {
	p := new(JsonProvider)
	dest := new(testStruct)
	err := p.Unmarshal([]byte(jsonExample1), dest)
	assert.NoError(t, err)
	assert.Equal(t, 1, dest.Test.Omg)
}
