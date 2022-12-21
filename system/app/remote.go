package app

import "github.com/valyala/fasthttp"

type Remote struct {
	req  *fasthttp.Request
	res  *fasthttp.Response
	body []byte
}

func (f *Remote) do() error {
	err := fasthttp.Do(f.req, f.res)
	fasthttp.ReleaseRequest(f.req)
	return err
}

func (f *Remote) Bytes() ([]byte, error) {
	if f.body != nil {
		return f.body, nil
	}

	f.do()

	f.body = f.res.Body()
	fasthttp.ReleaseResponse(f.res)

	return f.body, nil
}

func (b *Remote) String() (string, error) {
	data, err := b.Bytes()
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func NewRequest(url, methods string) *Remote {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(methods)
	req.SetRequestURI(url)

	res := fasthttp.AcquireResponse()

	return &Remote{
		req: req,
		res: res,
	}
}

func ReqGet(u string) *Remote {
	rm := NewRequest(u, "GET")
	return rm
}

func ReqPost(u string) *Remote {
	rm := NewRequest(u, "POST")
	return rm
}
