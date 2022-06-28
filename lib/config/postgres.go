package config

type Postgres struct {
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	DataLocation string	`yaml:"data_location"`
}