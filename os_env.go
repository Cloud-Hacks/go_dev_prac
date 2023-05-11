package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	envVarMap := map[string]string{}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		envVarMap[pair[0]] = pair[1]
	}

	fmt.Println(envVarMap["XMODIFIERS"])
}
