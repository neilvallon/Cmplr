package compiler

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

func JadeCompiler(f string) (b []byte, err error) {
	// Get file directory
	dir, err := filepath.Abs(f)
	if err != nil {
		return
	}

	// Open jade file
	file, err := ioutil.ReadFile(f)
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
		b = out
	} else {
		err = fmt.Errorf("%s\n%s", err, out)
	}

	return
}
