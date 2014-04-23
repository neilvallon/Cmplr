package project

import (
	"github.com/neilvallon/cmplr/compiler"
	"os"
	"path"
	"path/filepath"
	"fmt"
	"strings"
	"log"
)


type Job struct {
	Type string

	Outputfile string
	InputFiles []string

	Options map[string]string
}

func (j *Job) Run() {
	var err error
	defer func () {
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

func (j *Job) Watch() {
	if err := os.MkdirAll(path.Dir(j.Outputfile), 0777); err != nil {
		log.Println(err)
		return
	}

	c := compiler.New(j.InputFiles)

	out, err := c.CompileAsync();
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
