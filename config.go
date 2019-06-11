package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type ConfigInfo struct {
	ListenAddr string `toml:"listen_addr"`
	NoticeNick string `toml:"notice_nick"`
}

var Config ConfigInfo

//LoadConfig load configure
func LoadConfig(fp string) (err error) {
	println("config file  is ", fp)
	if _, err = toml.DecodeFile(fp, &Config); err != nil {
		fmt.Printf("Decode Config Error:%s \n", err.Error())
		return
	}

	fmt.Printf("Load Config file %s Success\n", fp)
	err = nil
	return
}
