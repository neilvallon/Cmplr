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

	serve := flag.Bool("s", false, "Start dev file server")
	watch := flag.Bool("w", false, "Watch files for changes")
	flag.Parse()

	for _, j := range cfg.Jobs {
		if *serve || *watch {
			go j.Watch()
		} else {
			j.Run()
		}
	}

	// File server
	if *serve {
		fmt.Println("Starting file server on port 8080")
		http.ListenAndServe(":8080", http.FileServer(http.Dir("./")))
	} else if *watch {
		<- make(chan bool) // Wait indefinitely
	}
}
