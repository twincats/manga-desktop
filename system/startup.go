package system

import (
	"context"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	WailsContext *context.Context
	arguments    []string
)

func Startup(ctx *context.Context) {
	WailsContext = ctx

	args := os.Args
	arguments = args

	runtime.EventsOnce(*WailsContext, "args", func(optionalData ...interface{}) {
		runtime.EventsEmit(*WailsContext, "args", arguments[1:])
	})

}
