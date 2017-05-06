package core

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
	"path/filepath"
)

type filter struct {
	directories map[string]byte
	extensions  map[string]byte
}

var inputFilter filter

// SetFilter set the filter dir and ext
func SetFilter(dirs []string, extensions []string) {
	inputFilter.directories = make(map[string]byte)
	inputFilter.extensions = make(map[string]byte)
	for _, dir := range dirs {
		inputFilter.directories[dir] = 0
	}
	for _, ext := range extensions {
		inputFilter.extensions[ext] = 0
	}
}

// OutputHash the output map with key:path, value:hash
var outputHash map[string]string

// GetOutputHash get the output hash value with path
func GetOutputHash() map[string]string {
	return outputHash
}

// calculateSHA2 don't check the input is file or not. It will assume input is file.
func calculateSHA2(filePath string) []byte {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return h.Sum(nil)
}

func isSkip(path string) bool {
	dir := filepath.Dir(path)
	if _, ok := inputFilter.directories[dir]; ok {
		return true
	}
	ext := filepath.Ext(path)
	if _, ok := inputFilter.extensions[ext]; ok {
		return true
	}
	return false
}

func walkCallback(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return filepath.SkipDir
	}
	if isSkip(path) {
		return nil
	}
	hash := calculateSHA2(path)
	outputHash[path] = string(hash)
	return nil
}

func getDevices() []string {
	var roots []string
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		path := string(drive) + ":\\"
		_, err := os.Open(path)
		if err == nil {
			roots = append(roots, path)
		}
	}
	return roots
}

// Start is the entry function will go through all hard disk
func Start() {
	roots := getDevices()
	for _, root := range roots {
		filepath.Walk(root, walkCallback)
	}
}
