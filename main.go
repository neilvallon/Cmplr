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

	fmt.Printf("Project: %s\n\n", cfg.ProjectName)
	
	for _, j := range cfg.Jobs {
		fmt.Println(j.Type)
		fmt.Println(j.Outputfile)
		fmt.Println(j.InputFiles)
		fmt.Printf("Options: %s\n\n", j.Options)
	}
}
