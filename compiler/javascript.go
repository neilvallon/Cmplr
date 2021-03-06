package compiler

import (
	"fmt"
	"os/exec"
)

func JsCompiler(f string) (b []byte, err error) {
	out, err := exec.Command("uglifyjs", f).CombinedOutput()
	if err == nil {
		b = out
	} else {
		err = fmt.Errorf("%s\n%s", err, out)
	}

	return
}
