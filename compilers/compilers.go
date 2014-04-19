package compilers

import (
	"path/filepath"
	"fmt"
)

type Compiler interface {
	Compile() error
	GetData() []byte
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