package core

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

const outputFileName = "result"

func writeToCSV() {
	pwd, _ := os.Getwd()
	outputPath := filepath.Join(pwd, outputFileName+".csv")
	f, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	csvWriter := csv.NewWriter(f)
	defer csvWriter.Flush()
	for key, value := range outputHash {
		err := csvWriter.Write(append([]string{key}, value...))
		if err != nil {
			fmt.Println(err)
		}
	}
}
