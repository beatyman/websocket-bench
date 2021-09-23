package impl

import (
	"context"
	"time"
)

type CommonAPI struct {
	//some other
}

func (n *CommonAPI) GetTime(ctx context.Context) (time.Time, error) {
	return time.Now(), nil
}
