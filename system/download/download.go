package download

import (
	"errors"
	"mangav4/system/download/types"
	"reflect"
)

type Download struct {
	types.Downloads
}

// Get new Download instance
func NewDownload() *Download {
	return new(Download)
}

func (f *Download) getDownloads(className string) types.Downloads {
	var d types.Downloads
	for c := range types.CLASS {
		if className == c {
			ref := reflect.ValueOf(types.CLASS[c])
			d = ref.Interface().(types.Downloads)
		}
	}
	return d
}

func (f *Download) GetChapter(o types.Option) (*types.Chapter, error) {
	if d := f.getDownloads(o.ServerName); d != nil {
		return d.GetChapter(o)
	}
	return nil, errors.New("error Server Name : " + o.ServerName + " Not Found or Implemented")
}

func (f *Download) GetPage(o types.Option) (*types.Page, error) {
	if d := f.getDownloads(o.ServerName); d != nil {
		return d.GetPage(o)
	}
	return nil, errors.New("error Server Name : " + o.ServerName + " Not Found or Implemented")
}
