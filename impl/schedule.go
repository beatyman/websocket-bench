package impl

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"
)

type worker struct {
	sealTasks chan<- interface{}
}

func (w *worker) run(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				w.sealTasks <- "ping from server"
			case <-ctx.Done():
				log.Warn("client exit : ")
			}

		}
	}()
}
