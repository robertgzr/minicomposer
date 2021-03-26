package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
)

func main() {
	filename := flag.String("file", "compose.yaml", "")
	flag.Parse()

	b, err := ioutil.ReadFile(*filename)
	if err != nil {
		panic(err)
	}
	config, err := loader.ParseYAML(b)
	if err != nil {
		panic(err)
	}
	composition, err := loader.Load(types.ConfigDetails{
		Version:     "",
		WorkingDir:  ".",
		ConfigFiles: []types.ConfigFile{{Filename: *filename, Config: config}},
		Environment: map[string]string{},
	})
	if err != nil {
		panic(err)
	}

	for _, s := range composition.Services {
		if s.Build != nil {
			panic("building not implemented")
		}
		if s.Image != "" {
			log.Printf("Pulling %s\n", s.Image)
			cmd := exec.Command("docker", "pull", s.Image)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				panic(err)
			}
		}
	}
}
