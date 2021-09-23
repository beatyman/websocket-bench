package impl

import (
	"context"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

type CommonAPI struct {
	//some other
}

func (n *CommonAPI) GetSession(context.Context) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (n *CommonAPI) WorkerQueue(ctx context.Context, cli interface{}) (<-chan interface{}, error) {
	ticker := time.NewTicker(time.Minute)
	queue := make(chan interface{})
	go func() {
		for {
			select {
			case <-ticker.C:
				queue <- "ping from server"
			case <-ctx.Done():
				log.Warn("client exit : ")
			}

		}
	}()
	return queue, nil
}

func (n *CommonAPI) WorkerDone(context.Context, interface{}) error {
	log.Info("client worker done ")
	return nil
}
