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

func (n *CommonAPI) WorkerQueue(ctx context.Context, id interface{}) (<-chan interface{}, error) {
	queue := make(chan interface{})
	worker := &worker{
		sealTasks: queue,
	}
	go worker.run(ctx)
	workerId, _ := id.(uuid.UUID)
	log.Info("new client connect to server ", workerId)
	return queue, nil
}

func (n *CommonAPI) WorkerDone(ctx context.Context, id interface{}) error {
	workerId, _ := id.(uuid.UUID)
	log.Info("client worker done a task ", workerId)
	return nil
}

func (n *CommonAPI) Version(ctx context.Context) (string, error) {
	return "v1.11.0", nil
}
