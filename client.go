package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cz-it/WeChatNotice/rpc"
	"google.golang.org/grpc"
)

// Client is a client cmd
type Client struct {
	svrAddr string
	conn    *grpc.ClientConn
	rpcCli  rpc.WeChatNoticeClient
}

// NewClient create a new client
func NewClient(addr string) (cli *Client) {
	cli = &Client{}
	cli.svrAddr = addr
	cli.init()
	return cli
}

func (cli *Client) init() {
	println("addr:", cli.svrAddr)
	conn, err := grpc.Dial(cli.svrAddr, grpc.WithInsecure())
	cli.conn = conn
	if err != nil {
		fmt.Printf("did not connect: %v \n", err)
		return
	}
	cli.rpcCli = rpc.NewWeChatNoticeClient(cli.conn)
}

func (cli *Client) notice(nick, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	req := &rpc.NoticeReq{
		Nick: nick,
		Msg:  msg,
	}
	defer cancel()
	rsp, err := cli.rpcCli.Notice(ctx, req)
	if err != nil {
		fmt.Printf("RPC Call Error:%s \n", err.Error())
	}
	fmt.Printf("Got Rsp:%v", rsp)
}
