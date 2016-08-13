package main

import (
	"fmt"
	"io"
	"os"
)

var (
	dir      = os.TempDir()
	filename = dir + "file"
)

func fail(s string, a ...interface{}) {
	fmt.Fprint(os.Stderr, fmt.Sprintf(s, a...)+"\n")
	os.Exit(2)
}

func main() {
	out, err := os.Create(filename)
	if err != nil {
		fail("cannot open file %s: %v", filename, err)
	}
	defer out.Close()

	_, err = io.Copy(out, os.Stdin)
	if err != nil {
		fail("cannot write file %s: %v", filename, err)
	}

	fmt.Fprint(os.Stdout, "Output to "+filename+"\n")
}
