package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func scanDirectory1(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path + "  1")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory1(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}
func reportPanic() {
	p := recover()
	if p == nil {
		return
	} else {
		err, ok := p.(error)
		if ok {
			fmt.Println(err)
		} else {
			panic(p)
		}
	}
}

func main() {
	defer reportPanic()
	//panic("some other issue")
	scanDirectory1("recover")

}
