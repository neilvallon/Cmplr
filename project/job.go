package project

import (
	"github.com/neilvallon/cmplr/compilers"
	"os"
	"path"
	"path/filepath"
	"fmt"
	"strings"
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

	var out []byte
	if out, err = compilers.Compile(j.InputFiles); err != nil {
		return
	}

	fo, err := os.Create(j.Outputfile)
	if err != nil {
		return
	}
	defer fo.Close()

	_, err = fo.Write(out)

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
