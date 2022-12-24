package internet

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"sync"
	"unsafe"
)

type ContentType string

func (c ContentType) String() string {
	return string(c)
}

const (
	FormType       ContentType = "application/x-www-form-urlencoded"
	FormUploadType ContentType = "multipart/form-data"
	JsonType       ContentType = "application/json"
	XmlType        ContentType = "application/xml"
	XmlTextType    ContentType = "text/xml"
)

// Client is Struct for making http request only need initialize once
type Client struct {
	sync.Mutex
	req *fasthttp.Request
	res *fasthttp.Response
	err *error
}

// getReq is getting *fasthttp.Request from pool
func (c *Client) getReq() {
	if c.req == nil {
		c.req = fasthttp.AcquireRequest()
		c.req.Header.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:108.0) Gecko/20100101 Firefox/108.0")
		c.req.Header.SetContentType(JsonType.String())
		c.req.Header.Set("Accept-Encoding", "gzip")
	}
}

// getRes is getting *fasthttp.Response from pool
func (c *Client) getRes() {
	if c.res == nil {
		c.res = fasthttp.AcquireResponse()
	}
}

// releaseReq Release Request to pool
func (c *Client) releaseReq() {
	fasthttp.ReleaseRequest(c.req)
}

// releaseRes Release Response to pool
func (c *Client) releaseRes() {
	fasthttp.ReleaseResponse(c.res)
}

// checkError is for checking client error when preparing request
func (c *Client) checkError() error {
	if c.err != nil {
		return *c.err
	}
	return nil
}

// do is for doing http Request
func (c *Client) do() error {
	c.getRes()
	return fasthttp.Do(c.req, c.res)
}

// SetContentType is setting Request Content-Type default : JsonType "application/json"
func (c *Client) SetContentType(contentType ContentType) *Client {
	c.getReq()
	c.req.Header.SetContentType(contentType.String())
	return c
}

// SetHeader is setting Request Header
func (c *Client) SetHeader(key string, value string) *Client {
	c.getReq()
	c.req.Header.Set(key, value)
	return c
}

// SetBody is for setting Request body Manually
func (c *Client) SetBody(body []byte) *Client {
	c.getReq()
	c.req.SetBody(body)
	return c
}

// SetPayload is for setting Payload to send with Post Request
func (c *Client) SetPayload(payload map[string]interface{}) *Client {
	c.getReq()
	stringPayload, err := json.Marshal(payload)
	if err != nil {
		c.err = &err
	}
	c.req.SetBody(stringPayload)
	return c
}

// RequestString is return Request string that being set
func (c *Client) RequestString() string {
	return c.req.String()
}

// Get create GET Request by URL
func (c *Client) Get(url string) *Client {
	c.Lock()
	defer c.Unlock()

	c.getReq()
	c.req.Header.SetMethod("GET")
	c.req.SetRequestURI(url)
	return c
}

// Post create POST Request by URL
func (c *Client) Post(url string) *Client {
	c.Lock()
	defer c.Unlock()

	c.getReq()
	c.req.Header.SetMethod("POST")
	c.req.SetRequestURI(url)
	return c
}

// Put create PUT Request by URL
func (c *Client) Put(url string) *Client {
	c.Lock()
	defer c.Unlock()

	c.getReq()
	c.req.Header.SetMethod("PUT")
	c.req.SetRequestURI(url)
	return c
}

// Delete create DELETE Request by URL
func (c *Client) Delete(url string) *Client {
	c.Lock()
	defer c.Unlock()

	c.getReq()
	c.req.Header.SetMethod("DELETE")
	c.req.SetRequestURI(url)
	return c
}

// Options create OPTIONS Request by URL
func (c *Client) Options(url string) *Client {
	c.Lock()
	defer c.Unlock()

	c.getReq()
	c.req.Header.SetMethod("OPTIONS")
	c.req.SetRequestURI(url)
	return c
}

// Bytes return response body in []byte value
func (c *Client) Bytes() ([]byte, error) {
	c.Lock()
	defer c.Unlock()

	// check error in preparation request
	if err := c.checkError(); err != nil {
		return nil, err
	}

	// check error in do calling request
	if err := c.do(); err != nil {
		return nil, err
	}

	defer c.releaseRes()
	defer c.releaseRes()

	contentType := c.res.Header.Peek("Content-Encoding")

	if bytes.EqualFold(contentType, []byte("gzip")) {
		return c.res.BodyGunzip()
	} else {
		return c.res.Body(), nil
	}

}

// String return response body in string value
func (c *Client) String() (string, error) {
	byteBody, err := c.Bytes()
	if err != nil {
		return "", err
	}
	return *(*string)(unsafe.Pointer(&byteBody)), nil
}

// BodyParser Parse Json []byte to any Pointer Struct type
func BodyParser[T any](data []byte, typeData *T) error {
	return json.Unmarshal(data, typeData)
}
