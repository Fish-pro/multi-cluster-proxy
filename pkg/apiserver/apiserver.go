package apiserver

import (
	"context"
	"fmt"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/apiserver/coreapi/cluster"
	clusterSvc "github.com/Fish-pro/multi-cluster-proxy/pkg/apiserver/coreapi/cluster/service"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/apiserver/coreapi/handler"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Fish-pro/multi-cluster-proxy/pkg/apiserver/middlewares"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/config"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/mlog"
)

const APIV1PathRoot = "/v1"

// APIServer aggregates all cube apis
type APIServer struct {
	*config.Config

	Server *http.Server
}

func registerCoreAPI(c *config.Config) http.Handler {
	router := gin.New()

	middlewares.SetUpMiddlewares(router)

	v1Router := router.Group(APIV1PathRoot, gin.BasicAuth(gin.Accounts{
		c.BasicAuthUser: c.BasicAuthPassword,
	}))

	router.GET("/ping", func(c *gin.Context) {
		mlog.Info("ping server success")
		c.JSON(200, map[string]string{
			"message": "ok",
		})
	})

	clusterService := clusterSvc.NewServiceContext(c)

	clusterMange := v1Router.Group("/clusters")
	{
		clusterMange.GET("", func(c *gin.Context) {
			cluster.ListClusters(c, clusterService)
		})
		clusterMange.GET("/:name", func(c *gin.Context) {
			cluster.GetCluster(c, clusterService)
		})
		clusterMange.POST("", func(c *gin.Context) {
			cluster.CreateCluster(c, clusterService)
		})
		clusterMange.PUT("/:name", func(c *gin.Context) {
			cluster.UpdateCluster(c, clusterService)
		})
		clusterMange.DELETE("/:name", func(c *gin.Context) {
			cluster.DeleteCluster(c, clusterService)
		})
	}

	k8sApiProxy := v1Router.Group("/proxy")
	{
		k8sApiProxy.Any("/clusters/:cluster/*url", func(c *gin.Context) {
			handler.ProxyHandle(c, clusterService)
		})
	}

	return router
}

func NewAPIServerWithOpts(c *config.Config) *APIServer {
	handler := registerCoreAPI(c)

	return &APIServer{
		Config: c,
		Server: &http.Server{
			Handler: handler,
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
