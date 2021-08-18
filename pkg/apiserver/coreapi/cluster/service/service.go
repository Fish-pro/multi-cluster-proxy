package service

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"

	"github.com/Fish-pro/multi-cluster-proxy/pkg/apiserver/coreapi/cluster/model"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/config"
)

type ServiceContext struct {
	Config *config.Config
	Model  model.ClusterModel // 手动代码
}

func NewServiceContext(c *config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewClusterModel(sqlx.NewMysql(c.DataSource)), // 手动代码
	}
}
