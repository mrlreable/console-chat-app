package cfg

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config types
type CfgType int

const (
	Yaml CfgType = iota
	Json
	Text
)

const rootPath = "."

// Default config values
var (
	cfgName string = "config"
	cfgPath string = rootPath
	cfgType string = ".yaml"
)

func SetConfigName(name string) {
	cfgName = name
}

func SetConfigPath(path string) {
	cfgPath = path
}

func SetConfigType(t CfgType) {
	switch t {
	case 0:
		cfgType = ".yaml"
	case 1:
		cfgType = ".json"
	case 2:
		cfgType = ".txt"
	default:
		panic(fmt.Sprintf("Unsupported config type %v", t))
	}
}

func NewConfig(out interface{}) error {
	root, err := getRoot()
	if err != nil {
		panic(err)
	}
	f, err := os.ReadFile(path.Join(root, cfgPath, cfgName+cfgType))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(f, out)
	if err != nil {
		return err
	}

	return nil
}

func getRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	root, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	return root, nil
}
