package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kechako/sjis2utf8/transform"
)

func _main() (int, error) {
	var err error

	var source io.Reader
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			return 1, err
		}
		defer file.Close()
		source = file
	} else {
		source = os.Stdin
	}

	reader := transform.NewReader(source)
	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		return 2, err
	}

	return 0, nil
}

func main() {
	exitstatus, err := _main()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] %v\n", err)
		os.Exit(exitstatus)
	}
}
