package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func scanDirectory(path string) error {
	fmt.Println(path)
	files,err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _,file := range files {
		filepath := filepath.Join(path,file.Name())
		if file.IsDir() {
			err := scanDirectory(filepath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(filepath)
		}
	}
	return nil
}

func main() {
	err := scanDirectory("recover")
	if err != nil {
		log.Fatal(err)
	}
}