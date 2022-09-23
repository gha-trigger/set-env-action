package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/sethvargo/go-githubactions"
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

	dataFile, err := os.Open(filepath.Join(dir, "data.yaml"))
	if err != nil {
		return err
	}
	defer dataFile.Close()
	data := map[string]interface{}{}
	if err := yaml.NewDecoder(dataFile).Decode(&data); err != nil {
		return err
	}

	eventFile, err := os.Open(filepath.Join(dir, "event.yaml"))
	if err != nil {
		eventFile, err = os.Open(filepath.Join(dir, "event.json"))
		if err != nil {
			return err
		}
	}
	defer eventFile.Close()
	event := map[string]interface{}{}
	if err := yaml.NewDecoder(eventFile).Decode(&event); err != nil {
		return err
	}
	data["event"] = event

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	githubactions.SetEnv("DATA", string(b))
	return nil
}
