package handle

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Fish-pro/multi-cluster-proxy/pkg/mlog"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/utils/errcode"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/utils/response"
)

var ts = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          50,
	IdleConnTimeout:       60 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
}

func ProxyHandle(c *gin.Context) {
	// http request params
	cluster := c.Param("cluster")
	url := c.Param("url")

	// get cluster info by cluster name
	host, _, _ := getClusterInfo(cluster)
	if host == "" {
		response.FailReturn(c, errcode.ClusterNotFoundError(cluster))
		return
	}
	// create director
	director := func(req *http.Request) {
		req.URL.Scheme = "https"
		req.URL.Host = host
		req.Host = host
		req.URL.Path = url
	}

	errorHandler := func(resp http.ResponseWriter, req *http.Request, err error) {
		if err != nil {
			mlog.Warn("cluster %s url %s proxy fail, %v", cluster, url, err)
			response.FailReturn(c, errcode.ServerErr)
			return
		}
	}

	requestProxy := &httputil.ReverseProxy{Director: director, Transport: ts, ErrorHandler: errorHandler}
	requestProxy.ServeHTTP(c.Writer, c.Request)
}

// get cluster info by clusterName
func getClusterInfo(clusterName string) (string, []byte, []byte) {
	mlog.Info("get cluster name: %s", clusterName)
	return "10.23.5.12:16443", []byte{}, []byte{}
}