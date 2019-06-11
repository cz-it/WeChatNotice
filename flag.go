package main

import (
	"flag"
	"fmt"
)

//Flag Flag
var Flag FlagInfo

//FlagInfo FlagInfo
type FlagInfo struct {
	Version    bool
	Daemon     bool
	Config     string
	CPUProfile string
	Notice     bool
	RPC        string
}

func init() {
	flag.Usage = func() {
		fmt.Printf("WeChat Notice\n")
		fmt.Println("Usage: WeChatNotice -[vdcn]")
		fmt.Println("	To publish a message use:")
		fmt.Println("		WeChatNotice -n nick message")
		flag.PrintDefaults()
	}

	flag.BoolVar(&Flag.Version, "v", false, "Show WeChatNotice's Version")
	flag.BoolVar(&Flag.Daemon, "d", false, "Start WeChatNotice as A Daemon")
	flag.StringVar(&Flag.Config, "c", "", "Config File of WeChatNotice")
	flag.StringVar(&Flag.RPC, "r", "", "RPC Addr of WeChatNotice")
	flag.BoolVar(&Flag.Notice, "n", false, "Use WeChatNotice as A Client to Publish a Notice")
	flag.Parse()

}
