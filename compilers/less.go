package compilers

import (
	"os/exec"
	"fmt"
)


func LessCompile(f string) (b []byte, err error) {
	out, err := exec.Command("lessc", f).CombinedOutput()
	if err == nil {
		b = out
	} else {
		err = fmt.Errorf("%s\n%s", err, out)
	}

	return
}
