package client

import (
	"context"
	"github.com/filecoin-project/go-jsonrpc"
	log "github.com/sirupsen/logrus"
	"net/http"
	"websocket-bench/api"
)



func NewCommonRPCV0(ctx context.Context, addr string, requestHeader http.Header) (api.Common, jsonrpc.ClientCloser, error) {
	var res api.CommonStruct
	log.Infof("%+v",api.GetInternalStructs(&res))
	closer, err := jsonrpc.NewMergeClient(ctx, addr, "Filecoin", api.GetInternalStructs(&res), requestHeader)
	return &res, closer, err
}
