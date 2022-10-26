package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func FileListCWD() []string {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var filelist []string

	for _, f := range files {
		filelist = append(filelist, f.Name())
	}

	return filelist
}

func FileListTREE() []string {
	var filelist []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		filelist = append(filelist, path)
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	return filelist
}
