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
func (cs CompilerSet) Compile() (out []byte, err error) {
	var bb [][]byte
	for _, c := range cs {
		if err = c.Compile(); err != nil {
			return
		}
		bb = append(bb, c.GetData())
	}

	out = bytes.Join(bb, []byte("\n"))

	return
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