package core

import (
	"github.com/go-liam/util/request"
	"grape/pkg/middleware/router"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var urlDefault = "/"

// handle :
func HandleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	// 跨域设置
	if router.AccessControlAllowHTTP(res, req) {
		return
	}
	if !auth(req) {
		io.WriteString(res, "bad ip")
		return
	}
	url := switchURL(req)
	println("url =", url)

	if url == "/" {
		Index(res, req)
		return
	}
	serveReverseProxy(url, res, req)
}

func switchURL(req *http.Request) string {
	path := req.URL.Path
	ls := strings.Split(path, "/")
	if path == "/" || ls == nil || len(ls) < 2 || ls[1] == "" {
		return urlDefault
	}
	for _, v := range list {
		if v.Path == ls[1] {
			return v.Site
		}
	}
	return urlDefault
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url2
	url2, _ := url.Parse(target)
	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url2)
	// Update the headers to allow for SSL redirection
	req.URL.Host = url2.Host
	req.URL.Scheme = url2.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Header.Set("X-Test-Host", "123.123.0.1")
	req.Host = url2.Host
	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

func Index(res http.ResponseWriter, req *http.Request) {
	ip := request.ClientIP(req)
	io.WriteString(res, "5-4-index,ip="+ip)
}
