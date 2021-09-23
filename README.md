# websocket-bench
模拟生产环境

### 编译
```
git clone https://github.com/beatyman/websocket-bench.git
cd websocket-bench
go build .

```

默认监听127.0.0.1:3500
### 服务端启动
```
./websocket-bench server --listen-addr=127.0.0.1:3500
```

###客户端启动
```
./websocket-bench  client --server-addr=ws://127.0.0.1:3500/rpc/v0

```