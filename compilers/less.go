package compilers

import (
	"os/exec"
)

type LessFile struct {
	Name  string
	Data    []byte

	Compress   bool
}

func (f *LessFile) Compile() (err error) {
	var comp string
	if f.Compress {
		comp = "--clean-css"
	}

	out, err := exec.Command("lessc", f.Name, comp).Output()
	if err == nil {
		f.Data = out
	}

	return
}

func (f *LessFile) GetData() []byte {
	return f.Data
}
