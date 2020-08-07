package metric

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"grape/configs/errorcode"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	stMethod = "method"
	stPath   = "path"
	stStatus = "status"
)

// SvMetric :
type SvMetric struct {
	IsOpen                bool // 开关
	metricHTTPReqDuration *prometheus.HistogramVec
	metricHTTPReqTotal    *prometheus.CounterVec
	IsRun                 bool
}

// Init ：
func (sv *SvMetric) Init() {
	if !sv.IsOpen {
		return
	}
	// 监控接口请求耗时
	// HistogramVec 是一组Histogram
	sv.metricHTTPReqDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "The HTTP request latencies in seconds.",
		Buckets: nil,
	}, []string{stMethod, stPath})
	// 这里的"stMethod"、"stPath" 都是label
	// 监控接口请求次数
	// HistogramVec 是一组Histogram
	sv.metricHTTPReqTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests made.",
	}, []string{stMethod, stPath, stStatus})
	// 这里的"method"、"path"、"status" 都是label
	prometheus.MustRegister(
		sv.metricHTTPReqDuration,
		sv.metricHTTPReqTotal,
	)
}

// /api/epgInfo/1371648200  -> /api/epgInfo
func (sv *SvMetric) parsePath(path string) string {
	max := 5 // max path : /api/v2/xxx/xx/
	itemList := strings.Split(path, "/")
	if len(itemList) > max {
		return strings.Join(itemList[0:max], "/")
	}
	return path
}

//NewMetric metric middleware
func (sv *SvMetric) NewMetric() gin.HandlerFunc {
	if !sv.IsOpen {
		return func(c *gin.Context) {
			c.Next()
		}
	}
	return func(c *gin.Context) {
		tBegin := time.Now()
		c.Next()
		duration := float64(time.Since(tBegin)) / float64(time.Second)
		path := sv.parsePath(c.Request.URL.Path)
		// 请求数加1
		sv.metricHTTPReqTotal.With(prometheus.Labels{
			stMethod: c.Request.Method,
			stPath:   path,
			stStatus: strconv.Itoa(c.Writer.Status()),
		}).Inc()

		//  记录本次请求处理时间
		sv.metricHTTPReqDuration.With(prometheus.Labels{
			stMethod: c.Request.Method,
			stPath:   path,
		}).Observe(duration)
	}
}

// Handler ：
func (sv *SvMetric) Handler() http.Handler {
	if !sv.IsOpen {
		return nil
	}
	return promhttp.Handler()
}

// HandlerClose ：
func (sv *SvMetric) HandlerClose(c *gin.Context) {
	c.String(errorcode.Int200, "is close")
}
