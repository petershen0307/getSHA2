package core

import (
	"path/filepath"
	"strings"
)

type filter struct {
	directories map[string]byte
	extensions  map[string]byte
}

// should ensure input path is not a directory
func (i filter) isSkip(path string) (bool, error) {
	dir := strings.ToLower(filepath.Dir(path))
	if _, ok := i.directories[dir]; ok {
		return true, filepath.SkipDir
	}
	ext := strings.ToLower(filepath.Ext(path))
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
		inputFilter.directories[strings.ToLower(filepath.Clean(dir))] = 0
	}
	for _, ext := range extensions {
		inputFilter.extensions[strings.ToLower(ext)] = 0
	}
	return inputFilter
}
