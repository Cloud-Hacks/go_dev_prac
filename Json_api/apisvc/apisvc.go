package apisvc

import (
	"context"
	"encoding/json"
	"net/http"
)

type Apisvr struct {
	svc Service.service
}

func NewApiReq(svc Service.service) Apisvr {
	return Apisvr{
		svc: svc,
	}
}

func (s *Apisvr) handleGetFact(w http.ResponseWriter, r *http.Request) {
	fact, err = s.svc.GetFactSvc(context.Background())

	if err != nil {

	}
}

func writeJson(w http.ResponseWriter, i int, t any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-type", "application/json")

	return json.NewEncoder(w).Encode(t)
	// return json.Marshal()
}
