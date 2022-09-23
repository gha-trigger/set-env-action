package main

import (
	"errors"
	"log"
	"os"

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
	failed := false
	for envName, expValue := range cfg.Envs {
		actValue := os.Getenv(envName)
		if expValue != actValue {
			log.Printf("[ERROR] the environment variable %s is wrong: wanted %s, got %s", envName, expValue, actValue)
			failed = true
		}
	}
	if failed {
		return errors.New("[ERROR] Some environment variables are wrong")
	}
	return nil
}
