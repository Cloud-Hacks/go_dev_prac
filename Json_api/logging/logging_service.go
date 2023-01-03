package logging

import (
	"context"
	"log"
	"time"

	"github.com/Cloud-Hacks/go_dev_prac/json_api/service"
)

type LoggingSvc struct {
	nxt service.Service
}

func (l *LoggingSvc) logGetfact(ctx context.Context) (f *service.myFact, err error) {
	defer func(strt time.Time) {
		log.Default("fact=%v, err=%v time taken=%v", f, err, time.Since(strt))
	}(time.Now())

	return &l.nxt.GetFactSvc(ctx)
}
