package limiter

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/request"
	"github.com/go-liam/util/response"
	"net/http"
	"strings"
)

/*
// Create a limiter struct.
    limiter := tollbooth.NewLimiter(1, time.Second, nil)
    r.GET("/", tollbooth_gin.LimitHandler(limiter), func(c *router.Context) {
        c.String(200, "Hello, world!")
    })
*/
const (
	errorRefuse = 801
)

// MaxNumber :
var MaxNumber float64 = 1000 // 1000

var errInfo *response.APIResponse
var ips []string

func init() {
	errInfo = response.APIResponseData(errorRefuse, "maximum request limit", nil)
	//ips = []string{"192.168.31.235","192.168.31.2","192.168.31.3"} //? 是否转 int 比较速度会快些？
}

// LimitHandler :
func LimitHandler(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			//c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.JSON(http.StatusOK, errInfo)
			c.Abort()
			//c.Status(402)
			//c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

// NewLimiter :
func NewLimiter(max float64) *limiter.Limiter {
	// Create a limiter struct.
	limiter := tollbooth.NewLimiter(max, nil)
	return limiter
}

// LimitHandlerIP :
func LimitHandlerIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := request.ClientIP(c.Request)
		//log.Println("ip:", ip)
		for _, v := range ips {
			if strings.Compare( v , ip) ==0 {
				c.Status(403)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
