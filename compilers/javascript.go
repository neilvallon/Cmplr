package compilers

import (
	"os/exec"
	"fmt"
)

type JsFile struct {
	Name  string
	Data  []byte
}

func (f *JsFile) Compile() (err error) {
	out, err := exec.Command("uglifyjs", f.Name).CombinedOutput()
	if err == nil {
		f.Data = out
	} else {
		err = fmt.Errorf("%s\n%s", err, out)
	}

	return
}

func (f *JsFile) GetData() []byte {
	return f.Data
}
