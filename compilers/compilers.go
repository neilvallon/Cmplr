package compilers

import (
	"path/filepath"
	"fmt"
	"bytes"
)

type Compiler interface {
	Compile() error
	GetData() []byte
}

type CompilerSet []Compiler
func (cs CompilerSet) Compile() (err error) {
	for _, c := range cs {
		if err = c.Compile(); err != nil {
			return
		}
	}
	return
}

func (cs CompilerSet) Output() []byte {
	var bb [][]byte
	for _, c := range cs {
		bb = append(bb, c.GetData())
	}

	return bytes.Join(bb, []byte("\n"))
}

func GetCompiler(file string) (c Compiler, err error) {
	switch ext := filepath.Ext(file); ext {
		case ".js":
			c = &JsFile{ Name: file }
		case ".less":
			c = &LessFile{ Name: file }
		case ".jade":
			c = &JadeFile{ Name: file }
		default:
			err = fmt.Errorf("Unsuported file type '%s'", ext)
	}
	return
}