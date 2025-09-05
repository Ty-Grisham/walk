package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func filterOut(path, ext string, minSize int64, info os.FileInfo) bool {
	// If path is a directory or the file size is less than the minimum
	// size specified by the user, return true
	if info.IsDir() || info.Size() < minSize {
		return true
	}

	// If the file's extension matches the extension specified by the user,
	// return true
	if ext != "" && filepath.Ext(path) != ext {
		return true
	}

	return false
}

func listFile(path string, out io.Writer) error {
	_, err := fmt.Fprintln(out, path)
	return err
}
