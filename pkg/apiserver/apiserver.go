package apiserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Fish-pro/multi-cluster-proxy/pkg/apiserver/coreapi/handle"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/apiserver/middlewares"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/config"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/mlog"
)

const APIPathRoot = "/v1"

// APIServer aggregates all cube apis
type APIServer struct {
	*config.Config

	Server *http.Server
}

func registerCoreAPI() http.Handler {
	router := gin.New()

	middlewares.SetUpMiddlewares(router)

	router.GET("/ping", func(c *gin.Context) {
		mlog.Info("ping server success")
		c.JSON(200, map[string]string{
			"message": "ok",
		})
	})

	k8sApiProxy := router.Group(APIPathRoot + "/proxy")
	{
		k8sApiProxy.Any("/clusters/:cluster/*url", handle.ProxyHandle)
	}

	return router
}

func NewAPIServerWithOpts(c *config.Config) *APIServer {
	router := registerCoreAPI()

	return &APIServer{
		Config: c,
		Server: &http.Server{
			Handler: router,
			Addr:    fmt.Sprintf("%s:%s", c.Host, c.Port),
		},
	}
}

func (s *APIServer) Initialize() error {
	return nil
}

func (s *APIServer) Run(stop <-chan struct{}) {
	go func() {
		var err error

		if s.Config.TlsCert != "" && s.Config.TlsKey != "" {
			err = s.Server.ListenAndServeTLS(s.Config.TlsCert, s.Config.TlsKey)
		} else {
			err = s.Server.ListenAndServe()
		}

		if err != nil {
			mlog.Fatal("core apiserver start err: %v", err)
		}
	}()

	mlog.Info("core apiserver listen in %s", s.Server.Addr)

	<-stop

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		mlog.Fatal("core apiserver forced to shutdown: %v", err)
	}
}
