package martiniglog

import (
	"log"
	"net/http"
	"time"

	"github.com/go-martini/martini"
	"github.com/golang/glog"
)

// Logger returns a middleware handler that logs the request as it goes in and the response as it goes out.
func Logger() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger) {
		addr := req.Header.Get("X-Real-IP")
		if addr == "" {
			addr = req.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = req.RemoteAddr
			}
		}

		rw := res.(martini.ResponseWriter)
		c.Next()

		now := time.Now().Format(time.RFC3339)
		glog.Infof("%s [%s] \"%s %s\" %d", addr, now, req.Method, req.URL.Path, rw.Status())
	}
}
