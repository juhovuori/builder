package server

import (
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

type Config interface {
	GetInt(key string) (error, int)
	GetString(key string) (error, string)
}

type BuilderConfig struct {
	ast *ast.File
}

func (cfg BuilderConfig) GetInt(key string) (error, int) {
	return nil, 1
}

func (cfg BuilderConfig) GetString(key string) (error, string) {
	return nil, ""
}

func ConfigFromString(in []byte) (Config, error) {
	parsed, err := hcl.ParseBytes(in)
	return BuilderConfig{parsed}, err
}

func DefaultConfig(filename string) (Config, error) {
	if filename == "" {
		filename = os.Getenv("BUILDER_CONFIG")
	}
	if filename == "" {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		filename = wd + "/builder.hcl"
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ConfigFromString(bytes)
}
