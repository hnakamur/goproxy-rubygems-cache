package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	docRoot := flag.String("root", ".", "document root directory")
	address := flag.String("http", ":8080", `HTTP service address (e.g., ":8080")`)
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose

	proxy.OnRequest(reqMethodIs("GET", "HEAD")).DoFunc(
		func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			filename := path.Join(*docRoot, ctx.Req.URL.Path)
			if !exists(filename) {
				return req, nil
			}

			bytes, err := ioutil.ReadFile(filename)
			if err != nil {
				ctx.Warnf("%s", err)
				return req, nil
			}
			resp := goproxy.NewResponse(req, "application/octet-stream",
				http.StatusOK, string(bytes))
			ctx.Logf("return response from local %s", filename)
			return req, resp
		})

	proxy.OnResponse(respReqMethodIs("GET", "HEAD")).Do(
		goproxy.HandleBytes(
			func(b []byte, ctx *goproxy.ProxyCtx) []byte {
				if ctx.Req.Method != "GET" || hasRespHeader(ctx.Resp, "Location") {
					return b
				}

				filename := path.Join(*docRoot, ctx.Req.URL.Path)
				if exists(filename) {
					return b
				}

				dir := path.Dir(filename)
				err := os.MkdirAll(dir, 0755)
				if err != nil {
					ctx.Warnf("cannot create directory: %s", dir)
				}

				err = ioutil.WriteFile(filename, b, 0644)
				if err != nil {
					ctx.Warnf("cannot write file: %s", filename)
				}

				ctx.Logf("save cache to %s", filename)

				return b
			}))
	log.Fatal(http.ListenAndServe(*address, proxy))
}

func reqMethodIs(methods ...string) goproxy.ReqConditionFunc {
	return func(req *http.Request, ctx *goproxy.ProxyCtx) bool {
		for _, method := range methods {
			if req.Method == method {
				return true
			}
		}
		return false
	}
}

func respReqMethodIs(methods ...string) goproxy.RespConditionFunc {
	return func(resp *http.Response, ctx *goproxy.ProxyCtx) bool {
		for _, method := range methods {
			if resp.Request.Method == method {
				return true
			}
		}
		return false
	}
}

func hasRespHeader(resp *http.Response, header string) bool {
	_, ok := resp.Header[header]
	return ok
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
