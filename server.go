package main

import (
	"context"
	"log"
	"net"

	"github.com/cz-it/WeChatNotice/rpc"
	"google.golang.org/grpc"
)

type Server struct {
	rpcSvr *grpc.Server
	addr   string
}

var server Server

// Notice is imp of gRPC's Notice
func (svr *Server) Notice(ctx context.Context, req *rpc.NoticeReq) (rsp *rpc.NoticeRsp, err error) {
	rsp = &rpc.NoticeRsp{
		Errno: 9527,
	}
	return rsp, nil
}

func (svr *Server) init(addr string) {
	svr.addr = addr
	svr.rpcSvr = grpc.NewServer()
	rpc.RegisterWeChatNoticeServer(svr.rpcSvr, svr)
}

func (svr *Server) start() {
	lis, err := net.Listen("tcp", svr.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := svr.rpcSvr.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
