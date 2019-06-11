package main

import (
	"flag"
	"fmt"

	"github.com/cz-it/serverkit/daemon"
)

func main() {

	InitLogger()

	if Flag.Version {
		fmt.Println("Cur Version:%s", "")
		return
	}

	if Flag.Config == "" && !Flag.Notice {
		flag.Usage()
		return
	}
	println("flag:", Flag.RPC)
	// For Client
	if Flag.Notice {
		fmt.Printf("args is %d \n", len(flag.Args()))
		if len(flag.Args()) != 2 {
			flag.Usage()
			return
		}
		nick := flag.Args()[0]
		msg := flag.Args()[1]
		cli := NewClient(Flag.RPC)
		cli.notice(nick, msg)
		return
	}

	// For Server
	if err := LoadConfig(Flag.Config); err != nil {
		fmt.Println("Loading Config Error")
		return
	}

	server.init(Config.ListenAddr)
	if Flag.Daemon {
		daemon.Boot("/tmp/wechat_notice.lock", "/tmp/wechat_notice.pid", func() {
			server.start()
		})
	} else {
		server.start()
	}
}
