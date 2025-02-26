package main

import (
	"fmt"
	"os"
)

func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return int(fi.Size()), nil
}

func main() {
	filenames := []string{
		"assets/file1.txt",
		"assets/file2.txt",
	}

	for _, filename := range filenames {
		len, err := fileLen(filename)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println(filename, "length is", len, "bytes")
	}
}
