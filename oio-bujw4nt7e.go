package main

import (
	"fmt"
	"log"
	"os"
)

// CreateClean creates a new directory and removes any existing files or directories with the same name
func CreateClean(dirName string) error {
	// Check if the directory already exists
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		// Create the directory if it does not exist
		if err := os.Mkdir(dirName, 0755); err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		// Remove the existing directory and its contents
		if err := os.RemoveAll(dirName); err != nil {
			return err
		}
		// Create the directory again
		if err := os.Mkdir(dirName, 0755); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := CreateClean("clean"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Directory created or cleaned successfully")
}