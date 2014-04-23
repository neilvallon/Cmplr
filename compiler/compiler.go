package compiler

import (
	"code.google.com/p/go.exp/fsnotify"
	"bytes"
	"log"
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

func (c *Compiler) Watch() chan []byte {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	ifUpdate := make(chan bool)
	go c.monitor(w, ifUpdate)

	for f := range c.files {
		err = w.WatchFlags(string(f), fsnotify.FSN_MODIFY)
		if err != nil {
			log.Fatal(err)
		}
	}

	ofUpdate := make(chan []byte)
	go func() {
		for {
			<-ifUpdate
			ofUpdate <- bytes.Join(c.cache, []byte{})
		}
	}()
	return ofUpdate
}

func (c *Compiler) monitor(w *fsnotify.Watcher, u chan bool) {
	for {
		select {
			case ev := <- w.Event:
				f := CmplrFile(ev.Name)
				if out, err := f.Compile(); err == nil {
					if !bytes.Equal(c.cache[c.files[f]], out) {
						log.Println(f, "- Recompiling")
						c.cache[c.files[f]] = out
						u <- true
					}
				} else {
					log.Println(err)
				}
			case err := <- w.Error:
				log.Println("error:", err)
		}
	}
}
