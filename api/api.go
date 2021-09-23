package api

import (
	"context"
	"golang.org/x/xerrors"
	"time"
)

var ErrNotSupported = xerrors.New("method not supported")

type Handler interface {
	GetTime(context.Context) (time.Time, error)
}

type Common interface {
	GetTime(context.Context) (time.Time, error)
}

type CommonStruct struct {
	Internal struct {
		GetTime func(context.Context) (time.Time, error)
	}
}
type CommonStub struct {
}

func (s *CommonStruct) GetTime(ctx context.Context) (time.Time, error) {
	if s.Internal.GetTime == nil {
		return *new(time.Time), ErrNotSupported
	}
	return s.Internal.GetTime(ctx)
}

func (s *CommonStub) GetTime(context.Context) (time.Time, error) {
	return *new(time.Time), ErrNotSupported
}
