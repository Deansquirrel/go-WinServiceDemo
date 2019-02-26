package worker

import (
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"time"
)

type worker struct {
}

func NewWorker() *worker {
	return &worker{}
}

func (p *worker) Do() {
	for {
		log.Info(goToolCommon.GetDateTimeStr(time.Now()))
		time.Sleep(time.Second * 5)
	}
}
