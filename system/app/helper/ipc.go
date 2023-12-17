package helper

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func NewIPC(c context.Context) *IPC {
	return &IPC{wailsContext: c}
}

type IPC struct {
	wailsContext context.Context
}

func (f *IPC) ListenEmit(eventName string, optionalData ...interface{}) {
	funcs, opData := FuncAndData(optionalData...)

	runtime.EventsOn(f.wailsContext, eventName, func(dataFromListen ...interface{}) {
		for _, fn := range funcs {
			opData = append(opData, fn(dataFromListen...))
		}

		runtime.EventsEmit(f.wailsContext, eventName+"_server", opData...)
	})
}

func (f *IPC) ListenEmitOnce(eventName string, optionalData ...interface{}) {
	funcs, opData := FuncAndData(optionalData...)

	runtime.EventsOnce(f.wailsContext, eventName, func(dataFromListen ...interface{}) {
		for _, fn := range funcs {
			opData = append(opData, fn(dataFromListen...))
		}

		runtime.EventsEmit(f.wailsContext, eventName+"_server", opData...)
	})
}

func (f *IPC) ListenOff(eventName string) {
	runtime.EventsOff(f.wailsContext, eventName)
}

func (f *IPC) ListenOffAll(eventName string) {
	runtime.EventsOffAll(f.wailsContext)
}

func (f *IPC) Emit(eventName string, optionalData ...interface{}) {
	runtime.EventsEmit(f.wailsContext, eventName, optionalData)
}

func (f *IPC) Listen(eventName string, callback func(optionalData ...interface{})) {
	runtime.EventsOn(f.wailsContext, eventName, callback)
}

func (f *IPC) ListenOnce(eventName string, callback func(optionalData ...interface{})) {
	runtime.EventsOnce(f.wailsContext, eventName, callback)
}

func (f *IPC) ListenTimes(eventName string, counter int, callback func(optionalData ...interface{})) {
	runtime.EventsOnMultiple(f.wailsContext, eventName, callback, counter)
}

// Function type with parameters as []interface{}
type FuncWithParams func(...interface{}) interface{}

/*
	FuncAndData is seperate array param interface{} and Func

Need return value if used you can return nil if not needed
*/
func FuncAndData(optionalData ...interface{}) ([]FuncWithParams, []interface{}) {
	var optionalFuncs []FuncWithParams
	var optionalOther []interface{}

	for _, data := range optionalData {
		switch v := data.(type) {
		case FuncWithParams:
			optionalFuncs = append(optionalFuncs, v)
		default:
			optionalOther = append(optionalOther, v)
		}
	}

	return optionalFuncs, optionalOther
}
