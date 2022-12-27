package main

import (
	"context"
	"net/http"
)

type Service interface {
	GetFactSvc(context.Context) (*myFact, error)
}

func GetFactSvc(context.Context) (*myFact, error) {
	resp, err := http.Get("http://catfact.ninja/fact")

	if err != nil {
		return nil, err
	}

	return resp, nil
}
