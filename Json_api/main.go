package main

import (
	"context"
	"fmt"

	"github.com/Cloud-Hacks/go_dev_prac/json_api/service"
)

func main() {
	ctx := context.Background()
	s := service.GetUrl("http://catfact.ninja/fact")

	f, _ := s.GetFactSvc(ctx)

	fmt.Println(f)

}
