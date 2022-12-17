package main

import (
	"context"
	"fmt"
	"time"
)

type statStack interface {
	pushfunc(string)
}

type ropt struct{}

type stat struct {
	curr string `json`
	stat int32  `json`
}

func (r ropt) pushfunc(ctx context.Context, temp string) string {
	ctx, canc := context.WithTimeout(ctx, time.Millisecond*500)
	defer canc()
	return temp
}

func main() {
	t := "uot"
	ctx := context.Background()
	r := &ropt{}

	f := r.pushfunc(ctx, t)

	fmt.Println(f)
}
