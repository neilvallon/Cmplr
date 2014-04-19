package project

import (
	"github.com/neilvallon/cmplr/compilers"
	"os"
	"path"
)


type Job struct {
	Type string

	Outputfile string
	InputFiles []string

	Options map[string]string
}

func (j *Job) Run() (err error) {
	if err = os.MkdirAll(path.Dir(j.Outputfile), 0777); err != nil {
		return
	}

	fo, err := os.Create(j.Outputfile)
	if err != nil {
		return
	}
	defer fo.Close()

	for _, f := range j.InputFiles {
		c := getCompiler(f)
		if err = c.Compile(); err == nil {
			_, err = fo.Write(c.GetData())
		}
	}

	return
}

func getCompiler(f string) compilers.Compiler {
	c, err := compilers.GetCompiler(f)
	if err != nil {
		panic(err)
	}

	return c
}
