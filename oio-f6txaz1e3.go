package main

import (
	"fmt"
	"log"
	"os"
)

// CreateClean creates a new directory and removes any existing files or directories with the same name
func CreateClean(dirName string) error {
	// Check if the directory already exists
	if err := os.RemoveAll(dirName); err != nil && !os.IsNotExist(err) {
		return err
	}

	// Create the directory
	if err := os.Mkdir(dirName, 0755); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := CreateClean("clean"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Directory created or cleaned successfully")
}