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

	var cl []compilers.Compiler
	for _, f := range j.InputFiles {
		c, err := compilers.GetCompiler(f)
		if err != nil {
			return err
		}
		cl = append(cl, c)
	}

	var bl [][]byte
	for _, c := range cl {
		if err = c.Compile(); err == nil {
			bl = append(bl, c.GetData())
		} else {
			return
		}
	}


	fo, err := os.Create(j.Outputfile)
	if err != nil {
		return
	}
	defer fo.Close()

	for _, b := range bl {
		_, err = fo.Write(b)
		if err != nil {
			return
		}
	}

	return
}
