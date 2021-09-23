package api

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

var ErrNotSupported = xerrors.New("method not supported")

type Common interface {
	GetSession(context.Context) (uuid.UUID, error)
	//测试长连接
	WorkerQueue(context.Context, interface{}) (<-chan interface{}, error)
	//测试任务提交
	WorkerDone(context.Context, interface{}) error
}

type CommonStruct struct {
	Internal struct {
		GetSession  func(context.Context) (uuid.UUID, error)
		WorkerQueue func(context.Context, interface{}) (<-chan interface{}, error)
		WorkerDone  func(context.Context, interface{}) error
	}
}

func (s *CommonStruct) GetSession(ctx context.Context) (uuid.UUID, error) {
	if s.Internal.GetSession == nil {
		return *new(uuid.UUID), ErrNotSupported
	}
	return s.Internal.GetSession(ctx)
}
func (s *CommonStruct) WorkerQueue(p0 context.Context, p1 interface{}) (<-chan interface{}, error) {
	if s.Internal.WorkerQueue == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.WorkerQueue(p0, p1)
}
func (s *CommonStruct) WorkerDone(p0 context.Context, p1 interface{}) error {
	if s.Internal.WorkerDone == nil {
		return ErrNotSupported
	}
	return s.Internal.WorkerDone(p0, p1)
}
