package system

import (
	"context"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var wailsContext *context.Context

var arguments []string

func Startup(ctx *context.Context) {
	wailsContext = ctx

	args := os.Args
	arguments = args

	runtime.EventsOnce(*wailsContext, "args", func(optionalData ...interface{}) {
		runtime.EventsEmit(*wailsContext, "args", arguments[1:])
	})

}
