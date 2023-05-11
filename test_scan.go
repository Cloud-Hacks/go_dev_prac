package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type desp interface {
	state_scan()
}

func state_scan() {

	text := "here i am newbie"

	// inp := os.Stdin
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		in := strings.ToLower(sc.Text())

		switch {
		// case in == "tesc":
		case strings.Contains(text, in):
			fmt.Println("Fine")
			return
		}
	}

}

func main() {
	state_scan()
}
