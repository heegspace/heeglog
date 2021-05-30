package heeglog

import (
	"context"
	"encoding/json"
	"time"

	"github.com/heegspace/heegproto/lognode"
	log "github.com/sirupsen/logrus"
	"github.com/asim/go-micro/v3"
)

var (
	local        bool
	gip          string
	gserver_name string
	s2sname      string
)

func Println(args ...interface{}) {
	log.Println(args)
}

func Json(obj interface{}) string {
	if nil == obj {
		return ""
	}

	data, _ := json.Marshal(obj)
	return string(data)
}

func Lognode(s2sname string, req *lognode.LogReq)  {
	if nil == req {
		return 
	}

	service := micro.NewService()
	service.Init()
	cli := lognode.NewLognodeService(s2sname, service.Client())
	cli.Log(context.Background(), req)

	return 
}

func Init(server_name, ip, _s2sname string, islocal bool) {
	gip = ip
	local = islocal
	gserver_name = server_name
	s2sname = _s2sname

	return
}

func Info(ctx context.Context, _func string, info string, extra map[string]string) {
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
		Ip:         gip,
		Extra:      extra,
	}

	Lognode(s2sname, req)

	return
}

func Debug(ctx context.Context, _func string, info string, extra map[string]string) {
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
		Ip:         gip,
		Extra:      extra,
	}

	Lognode(s2sname, req)

	return
}

func Warn(ctx context.Context, _func string, info string, extra map[string]string) {
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
		Ip:         gip,
		Extra:      extra,
	}

	Lognode(s2sname, req)

	return
}

func Error(ctx context.Context, _func string, info string, extra map[string]string) {
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
		Ip:         gip,
		Extra:      extra,
	}

	Lognode(s2sname, req)

	return
}

func CallInfo(ctx context.Context, _func string, req, res string, extra map[string]string) {
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
		Ip:         gip,
		Extra:      extra,
	}

	Lognode(s2sname, logreq)

	return
}

func CallDebug(ctx context.Context, _func string, req, res string, extra map[string]string) {
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
		Ip:         gip,
		Extra:      extra,
	}

	Lognode(s2sname, logreq)

	return
}

func CallWarn(ctx context.Context, _func string, req, res string, extra map[string]string) {
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
		Ip:         gip,
		Extra:      extra,
	}

	Lognode(s2sname, logreq)
	return
}

func CallError(ctx context.Context, _func string, req, res string, extra map[string]string) {
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
		Ip:         gip,
		Extra:      extra,
	}

	Lognode(s2sname, logreq)
	return
}
