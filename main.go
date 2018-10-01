package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	path := flag.String("path", "./", "path for files")
	withoutConfirm := flag.Bool("noconfirm", false, "delete without confirmation")
	flag.Parse()
	resultPath := "./"
	if len(*path) != 0 {
		resultPath = *path
	}
	if resultPath[len(resultPath)-1] != '/' {
		resultPath += "/"
	}
	files, err := ioutil.ReadDir(resultPath)
	if err != nil {
		panic(err)
	}
	hashes := make(map[string][]string, len(files))

	fmt.Printf("\tThe analysis of files is started ...\n\n")
	notify := make(chan result, 10)
	length := 0
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		filename := f.Name()
		length++
		go calcHash(resultPath, filename, notify)
	}
	index := 0
	for r := range notify {
		index++
		hashes[r.hash] = append(hashes[r.hash], r.filename)
		if index == length {
			close(notify)
		}
	}

	count := 0
	var copies = make([]string, 0)
	for _, files := range hashes {
		if len(files) > 1 {
			fmt.Printf("\tOriginal file %s, ", files[0])
			if len(files[1:]) > 1 {
				fmt.Printf("copies: %v\n", files[1:])
			} else {
				fmt.Printf("copy: %v\n", files[1])
			}
			copies = append(copies, files[1:]...)
			count++
		}
	}
	if count == 0 {
		fmt.Printf("\tNo duplicate files found in this folder\n")
		return
	}
	if !*withoutConfirm {
		fmt.Printf("\n\tFound %d duplicate files, delete copies? (originals will not be deleted) [y/N]: ", count)
		var v string
		if _, err := fmt.Scanln(&v); err != nil || !(v == "Y" || v == "y") {
			fmt.Printf("\tIf no, then no, exit the program\n")
			return
		}
	}

	successRemoved := 0
	for _, c := range copies {
		fmt.Printf("\tDelete file with name: %s\n", c)
		if err := os.Remove(resultPath + c); err != nil {
			fmt.Printf("\tCan't delete file with name: %s, error: %v\n", c, err)
			continue
		}
		successRemoved++
	}
	fmt.Printf("\t%d files successfully deleted!\n", successRemoved)
}

type result struct {
	filename string
	hash     string
}

func calcHash(path, filename string, notify chan<- result) {
	f, err := os.Open(path + filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		panic(err)
	}

	notify <- result{filename: filename, hash: string(h.Sum(nil))}
}
