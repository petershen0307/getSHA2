package core

import (
	"path/filepath"
	"reflect"
	"testing"
)

func Test_isSkip(t *testing.T) {
	type args struct {
		path string
	}
	SetFilter([]string{filepath.FromSlash("/a/b"), filepath.FromSlash("/aa/bb")}, []string{"exe"})
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"normal", args{filepath.FromSlash("/a/b/c.txt")}, true},
		{"normal", args{filepath.FromSlash("/za/zb/c.txt")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSkip(tt.args.path); got != tt.want {
				t.Errorf("isSkip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDevices(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
		{"list disks", []string{"C:\\"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDevices(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDevices() = %v, want %v", got, tt.want)
			}
		})
	}
}
