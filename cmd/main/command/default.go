package command

import (
	"flag"
	"github.com/pefish/go-commander"
	go_config "github.com/pefish/go-config"
	go_logger "github.com/pefish/go-logger"
	"net"
)

type DefaultCommand struct {

}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{

	}
}

func (dc *DefaultCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	flagSet.String("tcp-address", "0.0.0.0:8000", "<addr>:<port> to listen on for TCP clients")
	return nil
}

func (dc *DefaultCommand) OnExited() error {
	return nil
}

func (dc *DefaultCommand) Start(data commander.StartData) error {
	tcpAddress, err := go_config.Config.GetString("tcp-address")
	if err != nil {
		return err
	}
	tcpListener, err := net.Listen("tcp", tcpAddress)
	if err != nil {
		return err
	}
	go_logger.Logger.InfoF("listening on %s", tcpListener.Addr())

	select {}
}

