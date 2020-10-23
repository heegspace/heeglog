package heeglog

import (
	"context"
	"fmt"
	"time"

	"github.com/heegspace/heegproto/lognode"
	"github.com/heegspace/heegrpc"
	"github.com/heegspace/heegrpc/registry"
	"github.com/heegspace/heegrpc/rpc"
	log "github.com/sirupsen/logrus"
)

var (
	local        bool
	gip          string
	gserver_name string
	logClient    *rpc.HeegClient
)

func Init(logs2s *registry.S2sName, server_name, ip string, islocal bool) {
	if nil == logs2s {
		panic("logs2s info is nil!")
	}
	logClient = heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: logs2s.Host,
		Port: int(logs2s.Port),
	})

	gip = ip
	local = islocal
	gserver_name = server_name

	return
}

func Info(ctx context.Context, _func string, info string, extra map[string]string) {
	if nil == logClient {
		return
	}

	if local {
		log.Println("func: " + _func + "  info:" + info)

		return
	}

	req := &lognode.LogReq{
		Level:      lognode.LogLevel_INFO,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Func:       _func,
		Info:       info,
		ServerName: gserver_name,
		IP:         gip,
		Extra:      extra,
	}

	logNode := lognode.NewLognodeServiceClient(logClient.Client())
	err := logNode.Log(ctx, req)
	if nil != err {
		fmt.Println(req)
	}

	return
}

func Debug(ctx context.Context, _func string, info string, extra map[string]string) {
	if nil == logClient {
		return
	}

	if local {
		log.Debug("func: " + _func + "  info:" + info)

		return
	}

	req := &lognode.LogReq{
		Level:      lognode.LogLevel_INFO,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Func:       _func,
		Info:       info,
		ServerName: gserver_name,
		IP:         gip,
		Extra:      extra,
	}

	logNode := lognode.NewLognodeServiceClient(logClient.Client())
	err := logNode.Log(ctx, req)
	if nil != err {
		fmt.Println(req)
	}

	return
}

func Warn(ctx context.Context, _func string, info string, extra map[string]string) {
	if nil == logClient {
		return
	}

	if local {
		log.Warn("func: " + _func + "  info:" + info)

		return
	}

	req := &lognode.LogReq{
		Level:      lognode.LogLevel_INFO,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Func:       _func,
		Info:       info,
		ServerName: gserver_name,
		IP:         gip,
		Extra:      extra,
	}

	logNode := lognode.NewLognodeServiceClient(logClient.Client())
	err := logNode.Log(ctx, req)
	if nil != err {
		fmt.Println(req)
	}

	return
}

func Error(ctx context.Context, _func string, info string, extra map[string]string) {
	if nil == logClient {
		return
	}

	if local {
		log.Error("func: " + _func + "  info:" + info)

		return
	}

	req := &lognode.LogReq{
		Level:      lognode.LogLevel_INFO,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Func:       _func,
		Info:       info,
		ServerName: gserver_name,
		IP:         gip,
		Extra:      extra,
	}

	logNode := lognode.NewLognodeServiceClient(logClient.Client())
	err := logNode.Log(ctx, req)
	if nil != err {
		fmt.Println(req)
	}

	return
}

func CallInfo(ctx context.Context, _func string, req, res string, extra map[string]string) {
	if nil == logClient {
		return
	}

	if local {
		log.Info(_func + "  req: " + req + "  res: " + res)

		return
	}

	logreq := &lognode.CallLogReq{
		Level:      lognode.LogLevel_INFO,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Func:       _func,
		Req:        req,
		Res:        res,
		ServerName: gserver_name,
		IP:         gip,
		Extra:      extra,
	}

	logNode := lognode.NewLognodeServiceClient(logClient.Client())
	err := logNode.CallLog(ctx, logreq)
	if nil != err {
		fmt.Println(req)
	}

	return
}

func CallDebug(ctx context.Context, _func string, req, res string, extra map[string]string) {
	if nil == logClient {
		return
	}

	if local {
		log.Debug(_func + "  req: " + req + "  res: " + res)

		return
	}

	logreq := &lognode.CallLogReq{
		Level:      lognode.LogLevel_INFO,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Func:       _func,
		Req:        req,
		Res:        res,
		ServerName: gserver_name,
		IP:         gip,
		Extra:      extra,
	}

	logNode := lognode.NewLognodeServiceClient(logClient.Client())
	err := logNode.CallLog(ctx, logreq)
	if nil != err {
		fmt.Println(req)
	}

	return
}

func CallWarn(ctx context.Context, _func string, req, res string, extra map[string]string) {
	if nil == logClient {
		return
	}

	if local {
		log.Warn(_func + "  req: " + req + "  res: " + res)

		return
	}

	logreq := &lognode.CallLogReq{
		Level:      lognode.LogLevel_INFO,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Func:       _func,
		Req:        req,
		Res:        res,
		ServerName: gserver_name,
		IP:         gip,
		Extra:      extra,
	}

	logNode := lognode.NewLognodeServiceClient(logClient.Client())
	err := logNode.CallLog(ctx, logreq)
	if nil != err {
		fmt.Println(req)
	}

	return
}

func CallError(ctx context.Context, _func string, req, res string, extra map[string]string) {
	if nil == logClient {
		return
	}

	if local {
		log.Error(_func + "  req: " + req + "  res: " + res)

		return
	}

	logreq := &lognode.CallLogReq{
		Level:      lognode.LogLevel_INFO,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Func:       _func,
		Req:        req,
		Res:        res,
		ServerName: gserver_name,
		IP:         gip,
		Extra:      extra,
	}

	logNode := lognode.NewLognodeServiceClient(logClient.Client())
	err := logNode.CallLog(ctx, logreq)
	if nil != err {
		fmt.Println(req)
	}

	return
}
