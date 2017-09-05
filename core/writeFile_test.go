package core

import (
	"testing"
)

func Test_writeToCSV(t *testing.T) {
	outputHash = map[string][]string{"aaa": {"123", "zxc"}, "bbb": {"lkj", "654"}}
	t.Run("Test", func(t *testing.T) {
		writeToCSV()
	})
}
