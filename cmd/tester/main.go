package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func main() {
	if err := core(); err != nil {
		log.Fatal(err)
	}
}

func core() error {
	// read config file
	// test envs
	dir := os.Args[1]
	envs := map[string]string{}
	f, err := os.Open(filepath.Join(dir, "envs.yaml"))
	if err != nil {
		return err
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(&envs); err != nil {
		return err
	}
	failed := false
	for envName, expValue := range envs {
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
