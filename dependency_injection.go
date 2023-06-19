package main

import (
	"fmt"
	"time"
)

type Middleware interface {
	mdware()
}

type Midware struct {
	time time.Location
	md   Middleware
}

// type Fig struct{}

func (mid *Midware) mdware() {
	fmt.Println("mdware fetched")
	mid.time = *time.Now().UTC().Location()
}

func (mid *Midware) newMiddleware(mdl Middleware) *Midware {
	mid.md = mdl
	mid.md.mdware()
	// mdl.mdware()
	// mid.mdware()
	return &Midware{
		time: *time.FixedZone("Mumbai", 5),
		md:   mdl,
	}
}

func main() {
	md := &Midware{}
	// var mdi Middleware

	rs := md.newMiddleware(&Midware{})

	fmt.Println(rs)
}
