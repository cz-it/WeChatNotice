package main 

import (
	"fmt"
	"github.com/cz-it/serverkit/log"
)

//Logger is logcat used by all app
var Logger *log.Logger

func InitLogger() {
	var err error
	Logger, err = log.NewFileLogger("log", "WeChatNotice")
	if err != nil {
		fmt.Errorf("Create Logger Error\n")
		return
	}
	Logger.SetMaxFileSize(1024 * 1024 * 100) //100MB
	Logger.SetLevel(log.LDEBUG)
}