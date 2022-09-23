package main

import (
	"encoding/json"
	"log"
	"path/filepath"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	if err := core(); err != nil {
		log.Fatal(err)
	}
}

func core() error {
	files, err := filepath.Glob("testdata/*/*")
	if err != nil {
		return err
	}
	b, err := json.Marshal(files)
	if err != nil {
		return err
	}

	githubactions.SetEnv("FILES", string(b))
	return nil
}
