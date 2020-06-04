package servers

import (
	"go-mod/servers/client"
	"go-mod/servers/server"
)

type Service interface {
	Init(...Option)
	Options() Options
	Client() client.Client
	Server() server.Server
	Run() error
	String() string
}