package compilers

import (
	"os/exec"
	"io/ioutil"
	"fmt"
)

type JadeFile struct {
	Name  string
	Data  []byte
}

func (f *JadeFile) Compile() (err error) {
	// Open jade file
	file, err := ioutil.ReadFile(f.Name)
	if err != nil {
		return
	}

	// Run jade and wait on stdin
	cmd := exec.Command("jade", "--pretty")
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
