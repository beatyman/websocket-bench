package impl

import (
	"context"
	"github.com/google/uuid"
)

type CommonAPI struct {
	//some other
}

func (n *CommonAPI) GetSession(context.Context) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (n *CommonAPI) WorkerQueue(context.Context, interface{}) (<-chan interface{}, error) {
	return nil, nil
}

func (n *CommonAPI) WorkerDone(context.Context, interface{}) error {
	return nil
}
