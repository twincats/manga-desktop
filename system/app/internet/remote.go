package internet

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"sync"

	"mangav4/system/app/helper"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/imroc/req/v3"
)

var (
	ErrFinishDownload = errors.New("finish download")
)

type MdexPages struct {
	BaseUrl string
	Chapter struct {
		Hash      string
		Data      []string
		DataSaver []string
	}
}

// const (
// 	writeLimit = 2 * bwlimit.Mebibyte // write limit is 1048576 B/s
// 	readLimit  = 1 * bwlimit.Mebibyte // read limit is 4000 B/s
// )

type RemoteClient struct {
	*req.Client
}

func NewRemoteClient() *RemoteClient {
	c := &RemoteClient{
		Client: req.C().
			SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.44").
			SetCommonHeader("Accept-Encoding", "gzip"),
	}
	c.GetTransport().WrapRoundTripFunc(func(rt http.RoundTripper) req.HttpRoundTripFunc {
		return func(req *http.Request) (resp *http.Response, err error) {
			resp, err = rt.RoundTrip(req)
			if err != nil {
				return
			}
			encoding := resp.Header.Get("Content-Encoding")
			if encoding == "gzip" {
				gr, errs := gzip.NewReader(resp.Body)
				if errs != nil {
					err = errs
					return
				}
				resp.Body = io.NopCloser(gr)
			}
			return
		}
	})

	// dial := bwlimit.NewDialer(&net.Dialer{
	// 	Timeout:   30 * time.Second,
	// 	KeepAlive: 30 * time.Second,
	// }, writeLimit, readLimit)

	// c.SetDial(dial.DialContext)

	return c
}

func (c *RemoteClient) Download(u string, p string) error {
	resp := c.Get(u).SetOutputFile(p).Do()
	if resp.Err != nil {
		return resp.Err
	}
	return nil
}

func (c *RemoteClient) GETPages(pid string) (mdexPage MdexPages, err error) {
	u := fmt.Sprintf("https://api.mangadex.org/at-home/server/%s", pid)

	_, err = c.R().
		SetSuccessResult(&mdexPage).
		Get(u)

	return
}

type StatDownload struct {
	Index        int    `json:"index"`
	StatusResume bool   `json:"status_resume"`
	Filename     string `json:"filename"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	Err          error  `json:"err"`
}

func (f *StatDownload) SetError(e error) *StatDownload {
	f.Err = e
	return f
}

type FileDownload struct {
	c       *RemoteClient
	r       *req.Request
	counter int
	sync.Mutex
}

func NewFileDownload(client *RemoteClient) *FileDownload {
	return &FileDownload{
		c: client,
	}
}

func (f *FileDownload) autoDir(p string) error {
	return os.MkdirAll(p, os.ModePerm)
}

func (f *FileDownload) autoSetReq() {
	if f.r == nil {
		f.r = f.c.R()
	}
}

func (f *FileDownload) clearReq() {
	f.r = nil
}

func (f *FileDownload) SetReferer(ref string) *FileDownload {
	f.autoSetReq()
	f.r.SetHeader("Referer", ref)
	return f
}

func (f *FileDownload) SetHeader(key, val string) *FileDownload {
	f.autoSetReq()
	f.r.SetHeader(key, val)
	return f
}

// Download is to download and resume failed download with filename path and url
func (f *FileDownload) Download(filename, url string) StatDownload {
	f.Lock()
	defer f.Unlock()

	f.c.DisableAutoReadResponse()
	f.autoSetReq()

	var s StatDownload
	s.Filename = filename
	s.URL = url

	flag := os.O_CREATE | os.O_WRONLY
	whence := io.SeekStart

	//check if filename exist
	fs, err := os.Stat(filename)
	if err != nil {
		// return error except IsNotExist Error
		if !os.IsNotExist(err) {
			return *s.SetError(err)
		}
	}

	// check if file already exist, check complete file
	if fs != nil {
		// Do Req to get Header Content-length
		resp, err := f.c.R().Get(url)
		if err != nil {
			return *s.SetError(err)
		}

		contentLengthString := resp.Header.Get("Content-Length")
		// change contentLength "" to prevent error parseInt
		if contentLengthString == "" {
			contentLengthString = "0"
		}

		// convert content length to int64
		contentLength, err := strconv.ParseInt(contentLengthString, 10, 64)
		if err != nil {
			return *s.SetError(err)
		}

		// check header accept ranges bytes & Remote ContentLength > File Size
		if resp.Header.Get("Accept-Ranges") == "bytes" {
			if contentLength > fs.Size() {
				// set flag append and seek file to end
				flag = os.O_APPEND | os.O_WRONLY
				whence = io.SeekEnd

				// set Range Header to download only needed body
				f.r.SetHeader("Range", fmt.Sprintf("bytes=%d-", fs.Size()))
				s.StatusResume = true
			} else {
				// return error file already complete because content-length <= file size
				return *s.SetError(ErrFinishDownload)
			}

		}

		// close response from getting header
		resp.Body.Close()
	}

	// Open file to Create or Append data
	fi, err := os.OpenFile(filename, flag, 0666)
	if err != nil {
		return *s.SetError(err)
	}

	// Get Req Download
	f.autoSetReq()
	res, err := f.r.Get(url)
	if err != nil {
		return *s.SetError(err)
	}

	// set Seek to Start if create / seek to end if Append
	_, err = fi.Seek(0, whence)
	if err != nil {
		return *s.SetError(err)
	}

	// copy data from response body to file
	_, err = io.Copy(fi, res.Body)
	if err != nil {
		return *s.SetError(err)
	}

	// close response & file
	defer res.Body.Close()
	defer fi.Close()

	defer f.clearReq()

	// return status download
	return s

}

// DownloadBatch is to download files in batch with parallel download.
// b is workers, path is directory download, files is array of URLS & fn function each download.
func (f *FileDownload) DownloadBatch(b int, path string, files []string, fn func(s StatDownload)) []StatDownload {
	var o []StatDownload

	// set autoDir is create if dir is not created
	f.autoDir(path)

	// do parallel code
	helper.ParallelCodeI(b, files, func(i int, v string) {
		fname := fmt.Sprintf("%d%v", i+1, filepath.Ext(v))
		sd := f.Download(filepath.Join(path, fname), v)
		sd.Name = fname
		sd.Index = i

		// run fn function
		fn(sd)

		if sd.Err != nil {
			// Append ouput if only not ErrFinishDownload
			if sd.Err != ErrFinishDownload {
				o = append(o, sd)
			}
		}
	})

	// return only []Stat of Error in Download
	return o
}

// RetryDownloads is for Retrying Error Download []Stat nad return another failed download.
func (f *FileDownload) RetryDownloads(o []StatDownload, fn func(s StatDownload)) []StatDownload {
	var new_o []StatDownload

	helper.ParallelCodeI(1, o, func(i int, v StatDownload) {
		sd := f.Download(v.Filename, v.URL)
		sd.Name = v.Name
		sd.Filename = v.Filename
		sd.Index = v.Index

		fn(sd)
		if sd.Err != nil {
			if sd.Err != ErrFinishDownload {
				new_o = append(new_o, sd)
			}
		}
	})

	if len(new_o) > 0 && f.counter > 0 {
		fnew := f.RetryDownloads(new_o, fn)
		if len(fnew) > 0 {
			return fnew
		} else {
			return []StatDownload{}
		}
	}

	return []StatDownload{}
}

func (f *FileDownload) RetryDownloadsCounter(counter int, o []StatDownload, fn func(s StatDownload)) []StatDownload {
	f.counter = counter
	return f.RetryDownloads(o, fn)
}

// RetryDownload is for Retrying Error Download []Stat nad return another failed download.
func (f *FileDownload) RetryDownload(o []StatDownload, fn func(s StatDownload)) []StatDownload {
	var new_o []StatDownload
	helper.ParallelCodeI(1, o, func(i int, v StatDownload) {
		sd := f.Download(v.Filename, v.URL)
		fn(sd)
		if sd.Err != nil {
			if sd.Err != ErrFinishDownload {
				new_o = append(new_o, sd)
			}
		}
	})

	return new_o
}
