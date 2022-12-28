package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetFactSvc(context.Context) (*myFact, error)
}

type myFact struct{}

type factSvc struct {
	url string
}

func GetUrl(url string) Service {
	return &factSvc{
		url: url,
	}
}

func (s *factSvc) GetFactSvc(context.Context) (*myFact, error) {
	resp, err := http.Get("http://catfact.ninja/fact")
	resp.Body.Close()

	if err != nil {
		return nil, err
	}

	f := &myFact{}

	if err := json.NewDecoder(resp.Body).Decode(f); err != nil {
		return nil, err
	}

	return f, nil
}
