package router

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
)

// RouteTestTool :测试工具
func RouteTestTool(router *gin.Engine, method, url string, body io.Reader, utoken string, appSecret string) (string, error) {
	//router := SetupRouterTest()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	if utoken != "" {
		req.Header.Add("utoken", "jsrseida")
	}
	if appSecret != "" {
		req.Header.Add("appSecret", appSecret)
	}
	router.ServeHTTP(w, req)
	if http.StatusOK != w.Code {
		fmt.Printf("%v--%v", w.Code, url)
		return "", errors.New("httpStatusFail")
	}
	//可以先判断状态，再取值 assert.Equal(t, http.StatusOK, w.Code)
	body2 := w.Body.String()
	return body2, nil
}
