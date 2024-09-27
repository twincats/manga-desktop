package app

import (
	"context"
	"mangav4/system/app/helper"
	"mangav4/system/app/internet"
	"os"

	"gorm.io/gorm"
)

var (
	WailsContext *context.Context
	DB           *gorm.DB
	arguments    []string
	BlockClose   = false
	Client       = new(internet.Client)
	C            = internet.NewRemoteClient()
)

func Startup(ctx *context.Context) {
	WailsContext = ctx

	args := os.Args
	arguments = args

	// IPC event
	var ipc = helper.NewIPC(*WailsContext)
	ipc.ListenEmit("args", arguments[1:])

}
