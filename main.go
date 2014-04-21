package main

import (
	"github.com/neilvallon/cmplr/project"
	"fmt"
	"flag"
	"net/http"
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

	// File server
	serve := flag.Bool("s", false, "Start dev file server")
	flag.Parse()
	if *serve {
		http.ListenAndServe(":8080", http.FileServer(http.Dir("./")))
	}
}

