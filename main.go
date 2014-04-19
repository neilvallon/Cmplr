package main

import (
	"github.com/neilvallon/cmplr/project"
	"fmt"
	"strings"
	"path/filepath"
)

func main() {
	cfg, err := project.ReadConfig()
	if err != nil {
		panic("Error parsing config file.")
	}

	fmt.Printf("Project: %s\n", cfg.ProjectName)
	
	for _, j := range cfg.Jobs {
		PrintHeader(filepath.Base(j.Outputfile))
		if err := j.Run(); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Done.")
}

func PrintHeader(f string) {
	padwidth := 2
	if l := len(f); l < 78 {
		padwidth = 78 - l
	}

	pl := padwidth / 2
	pr := padwidth - pl
	fmt.Printf("%s %s %s\n", strings.Repeat("#", pl), f, strings.Repeat("#", pr))
}
