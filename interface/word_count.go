package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func LineCount(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// r implements *io.file
	var r io.Reader = file
	// io.ReadAll(r)

	// x,err :=io.Writer.Write(file)
	// io.WriteString()

	var lc LineCounter

	if strings.HasSuffix(fileName, ".gz") {
		r, err = gzip.NewReader(file)
		if err != nil {
			return 0, err
		}
	}

	if _, err := io.Copy(&lc, r); err != nil {
		return 0, err
	}

	return int(lc), nil
}

type LineCounter int

func (lc *LineCounter) Write(data []byte) (int, error) {
	for _, c := range data {
		if c == '\n' {
			*lc++
		}
	}

	return len(data), nil
}

func main() {
	fmt.Println(LineCount("testdata/road.txt"))
	fmt.Println(LineCount("testdata/aow.txt.gz"))
}
