package main

import (
	"github.com/pefish/go-binutils/cmd/main/command"
	"github.com/pefish/go-binutils/version"
	"github.com/pefish/go-commander"
	go_logger "github.com/pefish/go-logger"
)

func main() {
	commanderInstance := commander.NewCommander(version.AppName, version.Version, version.AppName + " 是一个二进制文件工具，祝你玩得开心。作者：pefish")
	//commanderInstance.RegisterSubcommand("client", client.NewClient())
	//commanderInstance.RegisterSubcommand("server", server.NewServer())
	commanderInstance.RegisterDefaultSubcommand(command.NewDefaultCommand())
	err := commanderInstance.Run()
	if err != nil {
		go_logger.Logger.Error(err)
	}
}
