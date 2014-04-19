package compilers

import (
	"os/exec"
	"io/ioutil"
	"path/filepath"
	"fmt"
)

type JadeFile struct {
	Name  string
	Data  []byte
}

func (f *JadeFile) Compile() (err error) {
	// Get file directory
	dir, err := filepath.Abs(f.Name)
	if err != nil {
		return
	}

	// Open jade file
	file, err := ioutil.ReadFile(f.Name)
	if err != nil {
		return
	}

	// Run jade and wait on stdin
	cmd := exec.Command("jade", "--pretty", "--path", dir)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return
	}

	// Pipe file into jade
	if _, err = stdin.Write(file); err != nil {
		return
	}
	stdin.Close()

	// Read output
	out, err := cmd.CombinedOutput()
	if err == nil {
		f.Data = out
	} else {
		err = fmt.Errorf("%s\n%s", err, out)
	}

	return
}

func (f *JadeFile) GetData() []byte {
	return f.Data
}
