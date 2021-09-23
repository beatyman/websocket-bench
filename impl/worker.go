package impl

import (
	"context"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

type worker struct {
	sealTasks chan<- int64
	taskCount int64
	ID        uuid.UUID
}

func (w *worker) run(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				w.taskCount++
				w.sealTasks <- w.taskCount
			case <-ctx.Done():
				log.Warn("client exit : ", w.ID)
				return
			}
		}
	}()
}
