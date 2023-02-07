package service

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetFactSvc(context.Context) (*myFact, error)
}

type myFact struct {
	Fact string `json:"fact"`
}

type factSvc struct {
	url string
}

func GetUrl(url string) Service {
	return &factSvc{
		url: url,
	}
}

func (s *factSvc) GetFactSvc(ctx context.Context) (*myFact, error) {
	resp, err := http.Get("http://catfact.ninja/fact")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	f := &myFact{}

	if err = json.NewDecoder(resp.Body).Decode(f); err != nil {
		return nil, err
	}

	return f, nil
}
