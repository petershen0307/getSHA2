package core

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type filter struct {
	directories map[string]byte
	extensions  map[string]byte
}

func (i filter) isSkip(path string) (bool, error) {
	dir := filepath.Dir(path)
	if _, ok := i.directories[dir]; ok {
		return true, filepath.SkipDir
	}
	ext := filepath.Ext(path)
	if _, ok := i.extensions[ext]; ok {
		return true, nil
	}
	return false, nil
}

// genFilter get the filter dir and ext
func genFilter(dirs []string, extensions []string) filter {
	var inputFilter filter
	inputFilter.directories = make(map[string]byte)
	inputFilter.extensions = make(map[string]byte)
	for _, dir := range dirs {
		inputFilter.directories[filepath.Dir(dir)] = 0
	}
	for _, ext := range extensions {
		inputFilter.extensions[ext] = 0
	}
	return inputFilter
}

// OutputHash the output map with key:hash, value:path list
var outputHash map[string][]string

// GetOutputHash get the output hash value with path
func GetOutputHash() map[string][]string {
	return outputHash
}

// calculateSHA2 don't check the input is file or not. It will assume input is file.
func calculateSHA2(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	sum := sha256.Sum256(content)
	return sum[:]
}

func genWalkCallback(i filter) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		if bSkip, ret := i.isSkip(path); bSkip {
			return ret
		}
		hash := calculateSHA2(path)
		hexStr := hex.EncodeToString(hash)
		outputHash[hexStr] = append(outputHash[hexStr], path)
		return nil
	}
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
func Start(path string, dirs, extensions []string) {
	roots := []string{path}
	if "" == path {
		roots = getDevices()
	}
	theFilter := genFilter(dirs, extensions)
	outputHash = make(map[string][]string)
	for _, root := range roots {
		filepath.Walk(root, genWalkCallback(theFilter))
	}
}
