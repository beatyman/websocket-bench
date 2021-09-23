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

func (n *CommonAPI) WorkerQueue(ctx context.Context, workerId uuid.UUID, last int64) (<-chan int64, error) {
	if last != 0 {
		log.Infof(" client  %v reconnect to server")
	} else {
		log.Info("new client connect to server ", workerId)
	}
	queue := make(chan int64)
	worker := &worker{
		sealTasks: queue,
		ID:        workerId,
		taskCount: last,
	}
	go worker.run(ctx)
	return queue, nil
}

func (n *CommonAPI) WorkerDone(ctx context.Context, workerId uuid.UUID, taskid int64) error {
	log.Infof("client worker %v done a task : %v", workerId, taskid)
	return nil
}

func (n *CommonAPI) Version(ctx context.Context) (string, error) {
	return "v1.11.0", nil
}
