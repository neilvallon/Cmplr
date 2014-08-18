package compiler

import (
	"bytes"
	"log"

	"gopkg.in/fsnotify.v0"
)

type Compiler struct {
	files map[CmplrFile]int // Filename to cache index
	cache [][]byte
}

// Returns new compiler from list of file paths.
func New(fs []string) *Compiler {
	// TODO: Error if file types do not match. For instance JS and coffeescript
	// compiled together makes sense but JS and less would not.
	c := &Compiler{
		files: make(map[CmplrFile]int),
		cache: make([][]byte, len(fs)),
	}

	for i, f := range fs {
		c.files[CmplrFile(f)] = i
	}

	return c
}

// Returns compiled source files appended together in the order given.
// Compilation will stop on first file error.
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

type retObj struct {
	id   int
	data []byte
	err  error
}

// Returns compiled source files appended together in the order given.
//
// Function will return on first error but compilation of all files will
// continue with their outputs discarded.
func (c *Compiler) CompileAsync() (b []byte, err error) {
	l := len(c.files)

	retchan := make(chan *retObj, l)
	for f, id := range c.files {
		go func(f CmplrFile, id int) {
			b, err := f.Compile()
			retchan <- &retObj{id: id, data: b, err: err}
		}(f, id)
	}

	for i := 0; i < l; i++ {
		ret := <-retchan
		if ret.err != nil {
			err = ret.err
			return
		}
		c.cache[ret.id] = ret.data
	}

	b = bytes.Join(c.cache, []byte{})
	return
}

// Runs compiler and watches file system for changes to any of the source files.
// When a modify event occurs that file is recompiled and combined with the
// cache of other files to be sent over a channel for saving.
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
		case ev := <-w.Event:
			if c.handelFileMod(ev) {
				u <- true
			}
		case err := <-w.Error:
			log.Println("error:", err)
		}
	}
}

func (c *Compiler) handelFileMod(ev *fsnotify.FileEvent) (up bool) {
	f := CmplrFile(ev.Name)

	out, err := f.Compile()
	if err != nil {
		log.Println(err)
		return
	}

	if !bytes.Equal(c.cache[c.files[f]], out) {
		log.Println(f, "- Recompiling")
		c.cache[c.files[f]] = out
		up = true
	}
	return
}
