package compilers

import (
	"os/exec"
)

type JsFile struct {
	Name  string
	Data  []byte
}

func (f *JsFile) Compile() (err error) {
	out, err := exec.Command("uglifyjs", f.Name).Output()
	if err == nil {
		f.Data = out
	}

	return
}

func (f *JsFile) GetData() []byte {
	return f.Data
}
