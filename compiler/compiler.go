package compiler

import (
	"bytes"
)

type Compiler struct {
	files map[CmplrFile]int // Filename to cache index
	cache [][]byte
}

func New(fs []string) *Compiler {
	c := &Compiler {
		files: make(map[CmplrFile]int),
		cache: make([][]byte, len(fs)),
	}

	for i, f := range fs {
		c.files[CmplrFile(f)] = i
	}

	return c
}

func (c *Compiler) Compile() (b []byte, err error) {
	for f, i := range c.files {
		b, err = f.Compile()
		if err != nil {
			return
		}
		c.cache[i] = b
	}

	b = bytes.Join(c.cache, []byte{})
	return
}
