package main

import (
	"log"

	"github.com/petershen0307/getSHA2/core"
)

func main() {
	core.SetFilter([]string{"C:\\Users\\PC\\Desktop\\code\\go\\src\\github.com\\petershen0307\\getSHA2\\.git"}, []string{})
	core.Start("C:\\Users\\PC\\Desktop\\code\\go\\src\\github.com\\petershen0307\\getSHA2")
	out := core.GetOutputHash()
	for key, value := range out {
		log.Println("key:", key, ", value:", value)
	}
}
