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
		if err := j.Run(); err != nil {
			fmt.Printf("\nCould not compile file: %s\n", j.Outputfile)
			fmt.Println(err)
		} else {
			fmt.Printf("\nSuccessfully compiled: %s\n", j.Outputfile)
		}
	}
}
