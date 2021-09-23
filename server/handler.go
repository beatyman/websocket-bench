package server

import (
	"context"
	"github.com/filecoin-project/go-jsonrpc"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"websocket-bench/impl"
)

func ServeRPC(ctx context.Context, listen string, maxRequestSize int64) error {
	rpcOpts := []jsonrpc.ServerOption{}
	if maxRequestSize > 0 {
		rpcOpts = append(rpcOpts, jsonrpc.WithMaxRequestSize(maxRequestSize))
	}

	rpcServer := jsonrpc.NewServer(rpcOpts...)
	rpcServer.Register("Filecoin", impl.NewCommonAPI(&impl.CommonAPI{}))

	http.Handle("/rpc/v0", rpcServer)

	server := http.Server{
		Addr:    listen,
		Handler: http.DefaultServeMux,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	sigCh := make(chan os.Signal, 2)

	go func() {
		select {
		case <-ctx.Done():

		case sig := <-sigCh:
			log.Infof("signal %s captured", sig)
		}

		if err := server.Shutdown(context.Background()); err != nil {
			log.Warnf("shutdown http server: %s", err)
		}

	}()

	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	log.Info("start http server", "addr", listen)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	log.Info("gracefull down")
	return nil
}
