package compilers

import (
	"os/exec"
	"fmt"
)

type LessFile struct {
	Name  string
	Data    []byte

	Compress   bool
}

func (f *LessFile) Compile(ch chan error) {
	var err error
	defer func() {
		ch <- err
	}()

	var comp string
	if f.Compress {
		comp = "--clean-css"
	}

	out, err := exec.Command("lessc", f.Name, comp).CombinedOutput()
	if err == nil {
		f.Data = out
	} else {
		err = fmt.Errorf("%s\n%s", err, out)
	}

	return
}

func (f *LessFile) GetData() []byte {
	return f.Data
}
