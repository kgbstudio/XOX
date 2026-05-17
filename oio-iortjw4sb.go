package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func clean(dir string) error {
	return filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err!= nil {
			return err
		}
		if entry.IsDir() && entry.Name() == "node_modules" {
			return os.RemoveAll(path)
		}
		if entry.Type().IsRegular() && (entry.Name() == "package-lock.json" || entry.Name() == "yarn.lock") {
			return os.Remove(path)
		}
		return nil
	})
}

func main() {
	if err := clean("."); err!= nil {
		log.Fatal(err)
	}
	fmt.Println("Cleaned")
}