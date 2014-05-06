package project

import (
	"fmt"
	"github.com/neilvallon/cmplr/compiler"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Job struct {
	Type string

	Outputfile string
	InputFiles []string

	Options map[string]string
}

// Compiles all files in current job asynchronously and saves to output file.
// If any error occurs durring compilation the job file will not be saved.
func (j *Job) Run() {
	var err error
	defer func() {
		printHeader(filepath.Base(j.Outputfile))
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err = os.MkdirAll(path.Dir(j.Outputfile), 0777); err != nil {
		return
	}

	c := compiler.New(j.InputFiles)

	var out []byte
	if out, err = c.CompileAsync(); err != nil {
		return
	}

	err = j.save(out)

	return
}

// Attempts an initial compile of all source files.
// If an error occurs on the inital compile a panic will occur to prevent a
// partial file from being saved.
//
// After the inital commit any changes that result in error will be logged and
// program will wait for the next successfull compile to continue saving.
func (j *Job) Watch() {
	if err := os.MkdirAll(path.Dir(j.Outputfile), 0777); err != nil {
		log.Println(err)
		return
	}

	c := compiler.New(j.InputFiles)

	// TODO: Fix panic by ensuring compiler cache is filled with all successful
	// data even if some files come back with errors.
	out, err := c.CompileAsync()
	if err != nil {
		panic(err) // invalid cache if initial compile fails
	}

	if err := j.save(out); err != nil {
		log.Println(err)
	}

	update := c.Watch()
	for out := range update {
		if err := j.save(out); err != nil {
			log.Println(err)
		}
	}
}

func (j *Job) save(b []byte) (err error) {
	fo, err := os.Create(j.Outputfile)
	if err != nil {
		return
	}
	defer fo.Close()

	_, err = fo.Write(b)

	return
}

func printHeader(f string) {
	padwidth := 2
	if l := len(f); l < 78 {
		padwidth = 78 - l
	}

	pl := padwidth / 2
	pr := padwidth - pl
	fmt.Printf("%s %s %s\n", strings.Repeat("#", pl), f, strings.Repeat("#", pr))
}
