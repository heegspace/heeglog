package heeglog

import (
	"time"
	"encoding/json"

	"github.com/heegspace/heegproto/lognode"
	"github.com/heegspace/heegrpc"
	"github.com/heegspace/heegrpc/registry"
	"github.com/heegspace/heegrpc/rpc"
)

var (
	logNode *lognode.LognodeServiceClient
)

func Init(s2s *registry.S2sName) {
	if nil == s2s {
		panic("s2s info is nil!")
	}
	client := heegrpc.NewHeegRpcClient(rpc.Option{
		Addr: s2s.Host,
		Port: int(s2s.Port),
	})

	logNode = lognode.NewLognodeServiceClientFactory(client.Client())

	return
}

func Info(ctx context.Context, _func string, info string, extra map[string,string]) {
	if nil == logNode {
		return 
	}

	req := &lognode.LogReq{
		Level:     lognode.LogLevel_INFO,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Func: _func,
		Info: info,
		Extra: extra,
	}

	logNode.Log(ctx, req)
	return 
}

func Debug(ctx context.Context, _func string, info string, extra map[string,string]) {
	if nil == logNode {
		return 
	}
	
	req := &lognode.LogReq{
		Level: lognode.LogLevel_INFO,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Func: _func,
		Info: info,
		Extra: extra,
	}

	logNode.Log(ctx, req)
	return 
}

func Warn(ctx context.Context, _func string, info string, extra map[string,string]) {
	if nil == logNode {
		return 
	}
	
	req := &lognode.LogReq{
		Level: lognode.LogLevel_INFO,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Func: _func,
		Info: info,
		Extra: extra,
	}

	logNode.Log(ctx, req)
	return 
}

func Error(ctx context.Context, _func string, info string, extra map[string,string]) {
	if nil == logNode {
		return 
	}
	
	req := &lognode.LogReq{
		Level: lognode.LogLevel_INFO,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Func: _func,
		Info: info,
		Extra: extra,
	}

	logNode.Log(ctx, req)
	return 
}

func CallInfo(ctx context.Context, _func string, req,res string, extra map[string,string]) {
	if nil == logNode {
		return 
	}
	
	req := &lognode.CallLogReq{
		Level: lognode.LogLevel_INFO,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Func: _func,
		Req: req,
		Res: res,
		Extra: extra,
	}

	logNode.CallLog(ctx, req)
	return 
}

func CallDebug(ctx context.Context, _func string, req,res string, extra map[string,string]) {
	if nil == logNode {
		return 
	}
	
	req := &lognode.CallLogReq{
		Level: lognode.LogLevel_INFO,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Func: _func,
		Req: req,
		Res: res,
		Extra: extra,
	}

	logNode.CallLog(ctx, req)
	return 
}

func CallWarn(ctx context.Context, _func string, req,res string, extra map[string,string]) {
	if nil == logNode {
		return 
	}
	
	req := &lognode.CallLogReq{
		Level: lognode.LogLevel_INFO,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Func: _func,
		Req: req,
		Res: res,
		Extra: extra,
	}

	logNode.CallLog(ctx, req)
	return 
}

func CallError(ctx context.Context, _func string, req,res string, extra map[string,string]) {
	if nil == logNode {
		return 
	}
	
	req := &lognode.CallLogReq{
		Level: lognode.LogLevel_INFO,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Func: _func,
		Req: req,
		Res: res,
		Extra: extra,
	}

	logNode.CallLog(ctx, req)
	return 
}
