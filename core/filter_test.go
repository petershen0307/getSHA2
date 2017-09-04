package core

import (
	"path/filepath"
	"testing"
)

func Test_isSkip(t *testing.T) {
	type args struct {
		path string
	}
	type want struct {
		bSkip bool
		ret   error
	}
	iFilter := genFilter([]string{filepath.FromSlash("/a/b"), filepath.FromSlash("/aa/bb")}, []string{"exe"})
	tests := []struct {
		name string
		args args
		want want
	}{
		// TODO: Add test cases.
		{"normal", args{filepath.FromSlash("/a/b/c.txt")}, want{true, nil}},
		{"normal", args{filepath.FromSlash("/za/zb/c.txt")}, want{false, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1, got2 := iFilter.isSkip(tt.args.path); got1 != tt.want.bSkip && got2 != tt.want.ret {
				t.Errorf("isSkip() = %v %v, want %v", got1, got2, tt.want)
			}
		})
	}
}
