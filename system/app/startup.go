package app

import (
	"context"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

var (
	WailsContext *context.Context
	DB           *gorm.DB
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
