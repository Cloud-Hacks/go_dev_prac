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
	curr string `json: "cur"`
	stat int32  `json: "stat`
}

func addStat() stat {
	return stat{
		curr: "yi",
		stat: 65,
	}
}

func (r ropt) pushfunc(ctx context.Context, temp string) string {
	ctx, canc := context.WithTimeout(ctx, time.Millisecond*500)
	defer canc()
	return temp
}

func (t stat) rpquest(tmp string) (string, int32) {
	t.curr = tmp
	return t.curr, t.stat
}

func main() {
	t := "uot"
	ctx := context.Background()

	s := addStat()
	x, y := s.rpquest(t)

	r := &ropt{}
	f := r.pushfunc(ctx, t)

	fmt.Println(f)
	fmt.Println(x, y)
}
