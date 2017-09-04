package core

import (
	"reflect"
	"testing"
)

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
