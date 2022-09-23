package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/sethvargo/go-githubactions"
	"gopkg.in/yaml.v3"
)

func main() {
	if err := core(); err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	Envs map[string]string
	Data interface{}
}

func core() error {
	// read config file
	// test envs
	file := os.Args[1]
	cfg := &Config{}
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
		return err
	}

	b, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	githubactions.SetEnv("DATA", string(b))
	return nil
}
