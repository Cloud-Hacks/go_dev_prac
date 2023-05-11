// input: 1  output: 1s
// input: 61 output: 1min1s
// input: 1000 output: ?h?min?s

package main

import (
	"fmt"
)

func main() {
	var (
		i       int32
		x, y, z int32
	)
	x, y, z = 0, 0, 0

	fmt.Scanf("%d", &i)

	// y := math.Remainder(i, 60)

	if i < 60 {
		z = i % 60
	} else if i >= 60 && i < 3600 {
		y = i / 60
		z = i % 60
	} else if i >= 3600 {
		x = i / 3600
		y = i % 3600
		if y > 60 {
			y = y / 60
			z = y % 60
		}
	}

	fmt.Printf("%dh: %dm: %ds", x, y, z)

}
