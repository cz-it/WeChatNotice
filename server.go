package main

import (
	"context"
	"fmt"
	"net"

	"github.com/cz-it/WeChatNotice/rpc"
	"github.com/cz-it/WeChatNotice/wxweb"
	"google.golang.org/grpc"
)

type Server struct {
	rpcSvr  *grpc.Server
	addr    string
	session *wxweb.Session
}

var server Server

// Notice is imp of gRPC's Notice
func (svr *Server) Notice(ctx context.Context, req *rpc.NoticeReq) (rsp *rpc.NoticeRsp, err error) {
	msg := &wxweb.TextMessage{}
	msg.Content = req.Msg
	msg.FromUserName = svr.session.Bot.UserName
	to := svr.session.Cm.GetContactByPYQuanPin(req.Nick)
	msg.ToUserName = to.UserName

	svr.session.SendChan <- msg

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

func msgHdl(session *wxweb.Session, msg *wxweb.ReceivedMessage) {
}

func (svr *Server) startWXWeb() {

	session, err := wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)

	if err != nil {
		fmt.Println(err)
		return
	}
	svr.session = session

	session.HandlerRegister.Add(wxweb.MSG_TEXT, wxweb.Handler(msgHdl), "msgHdl")
	session.HandlerRegister.EnableByName("msgHdl")
	go session.LoginAndServe(false)
}

func (svr *Server) start() {

	svr.startWXWeb()
	lis, err := net.Listen("tcp", svr.addr)
	if err != nil {
		fmt.Printf("failed to listen: %v \n", err)
	}

	fmt.Printf("Start RPC... \n")
	if err := svr.rpcSvr.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
