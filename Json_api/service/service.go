package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Service interface {
	GetFactSvc(context.Context) (*myFact, error)
}

type myFact struct {
	fact string `json`
}

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
	fmt.Println(resp.Body)

	if err != nil {
		return nil, err
	}

	f := &myFact{}

	if err = json.NewDecoder(resp.Body).Decode(f); err != nil {
		return nil, err
	}
	fmt.Println(f)

	return f, nil
}
