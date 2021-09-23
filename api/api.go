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
	WorkerQueue(context.Context, uuid.UUID, int64) (<-chan int64, error)
	//测试任务提交
	WorkerDone(context.Context, uuid.UUID, int64) error

	Version(context.Context) (string, error)
}

type CommonStruct struct {
	Internal struct {
		GetSession  func(context.Context) (uuid.UUID, error)
		WorkerQueue func(context.Context, uuid.UUID, int64) (<-chan int64, error)
		WorkerDone  func(context.Context, uuid.UUID, int64) error
		Version     func(context.Context) (string, error)
	}
}

func (s *CommonStruct) GetSession(ctx context.Context) (uuid.UUID, error) {
	if s.Internal.GetSession == nil {
		return *new(uuid.UUID), ErrNotSupported
	}
	return s.Internal.GetSession(ctx)
}
func (s *CommonStruct) WorkerQueue(p0 context.Context, p1 uuid.UUID, p2 int64) (<-chan int64, error) {
	if s.Internal.WorkerQueue == nil {
		return nil, ErrNotSupported
	}
	return s.Internal.WorkerQueue(p0, p1, p2)
}
func (s *CommonStruct) WorkerDone(p0 context.Context, p1 uuid.UUID, p3 int64) error {
	if s.Internal.WorkerDone == nil {
		return ErrNotSupported
	}
	return s.Internal.WorkerDone(p0, p1, p3)
}
func (s *CommonStruct) Version(p0 context.Context) (string, error) {
	if s.Internal.Version == nil {
		return "", ErrNotSupported
	}
	return s.Internal.Version(p0)
}
