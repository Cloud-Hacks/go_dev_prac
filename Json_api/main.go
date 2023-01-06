package main

import (
	"log"

	"github.com/Cloud-Hacks/go_dev_prac/json_api/apisvc"
	"github.com/Cloud-Hacks/go_dev_prac/json_api/service"
)

func main() {
	// ctx := context.Background()
	s := service.GetUrl("http://catfact.ninja/fact")

	// f, _ := s.GetFactSvc(ctx)

	apisvr := apisvc.NewApiReq(s)

	log.Fatal(apisvr.Handler(":5000"))

	// fmt.Println(*f)

}
