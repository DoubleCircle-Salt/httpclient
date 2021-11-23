package main

import "C"

import (
	"crypto/tls"

	"github.com/valyala/fasthttp"
)

const (
	VERSION = "0.0.1"
)

func main() {}

//export httpclient_version
func httpclient_version() *C.char {
	return C.CString(VERSION)
}

//export httpclient_post_with_sni
func httpclient_post_with_sni(url, host, sni, post_body *C.char) (code int, data *C.char) {
	
	reqClient := &fasthttp.HostClient{
		IsTLS:     true,
		Addr:      C.GoString(host),
		TLSConfig: &tls.Config{
			ServerName: C.GoString(sni),
		},
	}

	req := fasthttp.AcquireRequest()
	req.SetBody([]byte(C.GoString(post_body)))
	req.Header.SetMethod("POST")
	req.SetRequestURI(C.GoString(url))

	// parseURI
	req.URI()

	resp := fasthttp.AcquireResponse()

	err := reqClient.Do(req, resp)
	if err != nil {
		return -1, C.CString(err.Error())
	}

	fasthttp.ReleaseRequest(req)

	statusCode := resp.StatusCode()
	body := resp.Body()

	fasthttp.ReleaseResponse(resp)
	
	return statusCode, C.CString(string(body))
}
