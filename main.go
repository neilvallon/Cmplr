package main

import (
	"github.com/neilvallon/cmplr/project"
	"fmt"
)

func main() {
	cfg, err := project.ReadConfig()
	if err != nil {
		panic("Error parsing config file.")
	}

	fmt.Printf("Project: %s\n", cfg.ProjectName)

	for _, j := range cfg.Jobs {
		j.Run()
	}

	fmt.Println("Done.")
}

