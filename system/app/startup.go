package app

import (
	"context"
	"mangav4/system/app/internet"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

var (
	WailsContext *context.Context
	DB           *gorm.DB
	arguments    []string
	BlockClose   = false
	Client       = new(internet.Client)
	C            = new(internet.RemoteClient)
)

func Startup(ctx *context.Context) {
	WailsContext = ctx

	args := os.Args
	arguments = args

	runtime.EventsOnce(*WailsContext, "args", func(optionalData ...interface{}) {
		runtime.EventsEmit(*WailsContext, "args", arguments[1:])
	})

}
