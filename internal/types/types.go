package types

type Config struct {
	Address Address `yaml:"address"`
}

type Address struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
