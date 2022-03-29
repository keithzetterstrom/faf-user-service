package app

import (
	"flag"
	"os"

	db "github.com/keithzetterstrom/faf-user-service/internal/pkg/repository/postgres"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ServiceConfig  `yaml:"service"`
	PostgresConfig `yaml:"postgres"`
}

type ServiceConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type PostgresConfig struct {
	db.Config `yaml:"db"`
}

var (
	cnfPath = flag.String("config", "", "")
	netAddr = flag.String("addr", "", "")
)

func NewConfig(cfg *Config) error {
	flag.Parse()

	if *cnfPath == "" {
		return errors.New("no config path")
	}

	f, err := os.Open(*cnfPath)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	if *netAddr != "" {
		cfg.Host = *netAddr
	}

	return nil
}
