package impl

import (
	"context"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type CommonAPI struct {
	//some other
}

func (n *CommonAPI) GetSession(context.Context) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (n *CommonAPI) WorkerQueue(ctx context.Context, cli interface{}) (<-chan interface{}, error) {
	queue := make(chan interface{})
	worker := &worker{
		sealTasks: queue,
	}
	go worker.run(ctx)
	return queue, nil
}

func (n *CommonAPI) WorkerDone(context.Context, interface{}) error {
	log.Info("client worker done ")
	return nil
}
