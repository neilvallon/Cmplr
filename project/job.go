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

	cl, err := j.makeCompilers()
	if err != nil {
		return
	}

	if err = cl.Compile(); err != nil {
		return
	}
	out := cl.Output()

	fo, err := os.Create(j.Outputfile)
	if err != nil {
		return
	}
	defer fo.Close()

	_, err = fo.Write(out)

	return
}

func (j *Job) makeCompilers() (cl compilers.CompilerSet, err error) {
	if err = os.MkdirAll(path.Dir(j.Outputfile), 0777); err != nil {
		return
	}

	for _, f := range j.InputFiles {
		c, e := compilers.GetCompiler(f)
		if e != nil {
			err = e
			break
		}
		cl = append(cl, c)
	}
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
