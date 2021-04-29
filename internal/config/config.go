package config

import (
	"io/ioutil"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/daangn/guestbook/cmd/guestbook/cli"
)

type Config struct {
	Listen string `yaml:"listen"`
	GRPC   GRPC   `yaml:"grpc"`
	MySQL  MySQL  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
}

type GRPC struct {
	MaxConnectionIdle     duration `yaml:"max_connection_idle"`
	MaxConnectionAge      duration `yaml:"max_connection_age"`
	MaxConnectionAgeGrace duration `yaml:"max_connection_age_grace"`
	Time                  duration `yaml:"time"`
	Timeout               duration `yaml:"timeout"`
}

type MySQL struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

func NewConfig(flags *cli.Flags) (*Config, error) {
	return parse(flags.ConfigPath)
}

func parse(configPath string) (*Config, error) {
	filename, _ := filepath.Abs(configPath)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
