package server

import (
	"github.com/go-liam/util/request/tshttp"
	"strings"
	"testing"
	"time"
)

const (
	tsHTTPPort    = ":4402"
	tsHTTPAddress = "http://localhost" + tsHTTPPort
)

func TestHTTPIndex(t *testing.T) {
	go RunServerGin(tsHTTPPort)
	time.Sleep(time.Microsecond * 10)
	tsIndex()
}

func tsIndex()  {
	sv := new(tshttp.HTTPTest)
	url := tsHTTPAddress+"/"
	data := `
		{
			"ID":1,
			"versionId":101
		}
		`
	reader := strings.NewReader(data)
	code, back, _ := sv.HTTPData(tshttp.MethodHTTPGET,url,reader,"","")
	println("code:",code)
	println(back)
}
