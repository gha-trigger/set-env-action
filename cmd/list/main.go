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
	envFiles, err := filepath.Glob("testdata/*/*/envs.yaml")
	if err != nil {
		return err
	}
	dirs := make([]string, len(envFiles))
	for i, envFile := range envFiles {
		dirs[i] = filepath.Dir(envFile)
	}

	b, err := json.Marshal(dirs)
	if err != nil {
		return err
	}

	githubactions.SetEnv("DIRS", string(b))
	return nil
}
