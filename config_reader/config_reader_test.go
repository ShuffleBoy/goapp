package config_reader

import (
	"errors"
	"github.com/ShuffleBoy/goapp/config_reader/mocks"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_getFileContents(t *testing.T) {
	mocker := func() IoHelper {
		m := new(mocks.IoHelper)
		m.On("ReadFile", "file1").Return([]byte("some content"), nil)
		m.On("ReadFile", "file2").Return(nil, errors.New("wow"))
		return m
	}()

	type args struct {
		filePath string
		helper   IoHelper
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "all ok",
			args: args{
				filePath: "file1",
				helper:   mocker,
			},
			want: []byte("some content"),
		},
		{
			name: "with err",
			args: args{
				filePath: "file2",
				helper:   mocker,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFileContents(tt.args.filePath, tt.args.helper)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFileContents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFileContents() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadConfig(t *testing.T) {
	mocker := func() IoHelper {
		m := new(mocks.IoHelper)
		m.On("ReadFile", "json1").Return([]byte(jsonExample1), nil)
		m.On("ReadFile", "yaml1").Return([]byte(yamlExample1), nil)
		return m
	}()
	expValue := &testStruct{Test: struct {
		Omg int `yaml:"omg"`
	}(struct{ Omg int }{Omg: 1})}
	type args struct {
		filePath string
		dest     any
		provider Provider
		helper   IoHelper
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				filePath: "json1",
				dest:     new(testStruct),
				provider: new(JsonProvider),
				helper:   mocker,
			},
		},
		{
			args: args{
				filePath: "yaml1",
				dest:     new(testStruct),
				provider: new(YamlProvider),
				helper:   mocker,
			},
		},
		{
			args: args{
				filePath: "yaml1",
				dest:     new(testStruct),
				provider: new(JsonProvider),
				helper:   mocker,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadConfig(tt.args.filePath, tt.args.dest, tt.args.provider, tt.args.helper); (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				assert.Equal(t, expValue, tt.args.dest)
			}
		})
	}
}
