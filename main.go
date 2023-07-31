package main

import (
	"path/filepath"

	"sugilite.frank-mayer.io/parser"
)

func main() {
    f, err := filepath.Abs("./sample/main.sugilite")
    if err != nil {
        panic(err)
    }
    ast := parser.Parse(f)
}
