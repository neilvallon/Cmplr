package compiler

import (
	"fmt"
	"path/filepath"
)

type CmplrFile string
type CompileFunc func(string) ([]byte, error)

func (cf *CmplrFile) Compile() (b []byte, err error) {
	f, err := cf.getCompiler()
	if err == nil{
		b, err = f(string(*cf))
	}
	return
}

func (cf *CmplrFile) getCompiler() (c CompileFunc, err error) {
	switch ext := cf.ext(); ext {
		case ".js":
			c = JsCompiler
		case ".less":
			c = LessCompiler
		case ".jade":
			c = JadeCompiler
		default:
			err = fmt.Errorf("Unsuported file type '%s'", ext)
	}
	return
}

func (cf CmplrFile) ext() string {
	return filepath.Ext(string(cf));
}
