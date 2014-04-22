package compilers

import (
	"path/filepath"
	"fmt"
	"bytes"
)


type Compiler func(string) ([]byte, error)

type CompileOut struct {
	Id    int
	Data  []byte
	Error error
}

func Compile(fs []string) (b []byte, err error) {
	l := len(fs)

	retchan := make(chan *CompileOut, l)
	for i, f := range fs {
		var c Compiler
		if c, err = GetCompiler(f); err != nil {
			return
		}

		go func(f string, i int) {
			out, err := c(f)
			retchan <- &CompileOut{ Id: i, Data: out, Error: err, }
		}(f, i)
	}

	bb := make([][]byte, l)
	for i := 0; i < l; i++ {
		if ret := <- retchan; ret.Error == nil {
			bb[ret.Id] = ret.Data
		} else {
			err = ret.Error
			return
		}
	}

	b = bytes.Join(bb, []byte("\n"))
	return
}

func GetCompiler(file string) (c Compiler, err error) {
	switch ext := filepath.Ext(file); ext {
		case ".js":
			c = JsCompile
		case ".less":
			c = LessCompile
		case ".jade":
			c = JadeCompile
		default:
			err = fmt.Errorf("Unsuported file type '%s'", ext)
	}
	return
}