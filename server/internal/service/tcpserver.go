// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/api/servmsg"
	"hotgo/internal/library/cron"
	"hotgo/internal/library/network/tcp"
)

type (
	ITCPServer interface {
		OnAuthSummary(ctx context.Context, req *servmsg.AuthSummaryReq)
		CronDelete(ctx context.Context, in *servmsg.CronDeleteReq) (err error)
		CronEdit(ctx context.Context, in *servmsg.CronEditReq) (err error)
		CronStatus(ctx context.Context, in *servmsg.CronStatusReq) (err error)
		CronOnlineExec(ctx context.Context, in *servmsg.CronOnlineExecReq) (err error)
		DispatchLog(ctx context.Context, in *servmsg.CronDispatchLogReq) (log *cron.Log, err error)
		OnExampleHello(ctx context.Context, req *servmsg.ExampleHelloReq)
		OnExampleRPCHello(ctx context.Context, req *servmsg.ExampleRPCHelloReq) (res *servmsg.ExampleRPCHelloRes, err error)
		Instance() *tcp.Server
		Start(ctx context.Context)
		Stop(ctx context.Context)
		DefaultInterceptor(ctx context.Context, msg *tcp.Message) (err error)
		PreFilterInterceptor(ctx context.Context, msg *tcp.Message) (err error)
	}
)

var (
	localTCPServer ITCPServer
)

func TCPServer() ITCPServer {
	if localTCPServer == nil {
		panic("implement not found for interface ITCPServer, forgot register?")
	}
	return localTCPServer
}

func RegisterTCPServer(i ITCPServer) {
	localTCPServer = i
}
