package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

// Remember of separation of concerns!
// The next method are breaking the Single-responsibility principle!
// The Journal struct should be only responsible to manage Journal notes.
// Save method should be handled by another broader interface specialized in
// file management.

func (j *Journal) Save(filename string) {
	content := strings.Join(j.entries, "\n")
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		panic("Error to save journal.")
	}
}

// Example of a separed interface to save files
type FileManager struct {
	fileName string
}

func (f *FileManager) Save(content string) {
	err := ioutil.WriteFile(f.fileName, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	j := Journal{}
	j.AddEntry("Today I'm feeling great!")
	j.AddEntry("SOLID is cool!")

	f := FileManager{"jornal.txt"}
	content := strings.Join(j.entries, "\n")
	f.Save(content)
}
