package apisvc

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Cloud-Hacks/go_dev_prac/json_api/service"
)

type Apisvr struct {
	svc service.FactSvc
}

func NewApiReq(svc service.FactSvc) Apisvr {
	return Apisvr{
		svc: svc,
	}
}

func (s *Apisvr) Handler(listener string) error {
	http.HandleFunc("/", s.handleGetFact)
	return http.ListenAndServe(listener, nil)
}

func (s *Apisvr) handleGetFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.GetFactSvc(context.Background())

	if err != nil {
		writeJson(w, http.StatusExpectationFailed, map[string]any{"error": err.Error()})
		return
	}

	writeJson(w, http.StatusAccepted, fact)
}

func writeJson(w http.ResponseWriter, i int, t any) error {
	w.WriteHeader(i)
	w.Header().Add("Content-type", "application/json")

	return json.NewEncoder(w).Encode(t)
	// return json.Marshal()
}
