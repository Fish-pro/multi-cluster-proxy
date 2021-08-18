package cluster

import (
	"github.com/gin-gonic/gin"

	"github.com/Fish-pro/multi-cluster-proxy/pkg/apiserver/coreapi/cluster/model"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/apiserver/coreapi/cluster/service"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/mlog"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/utils/errcode"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/utils/response"
)

func ListClusters(c *gin.Context, svc *service.ServiceContext) {

	clusters, err := svc.Model.List()
	if err != nil {
		mlog.Error("failed to list cluster err: %v", err.Error())
		response.FailReturn(c, errcode.GetResourceError("cluster"))
		return
	}

	response.SuccessReturn(c, clusters)
	return
}

func GetCluster(c *gin.Context, svc *service.ServiceContext) {

	name := c.Param("name")

	cluster, err := svc.Model.FindOne(name)
	if err != nil {
		mlog.Error("failed to get cluster %s err: %v", name, err.Error())
		response.FailReturn(c, errcode.ClusterNotFoundError(name))
		return
	}

	response.SuccessReturn(c, cluster)
	return
}

func CreateCluster(c *gin.Context, svc *service.ServiceContext) {

	clusterRequest := model.Cluster{}
	err := c.BindJSON(&clusterRequest)
	if err != nil {
		mlog.Error("get create cluster body err:%s", err.Error())
		response.FailReturn(c, errcode.CreateResourceError("cluster"))
		return
	}

	_, err = svc.Model.Insert(clusterRequest)
	if err != nil {
		mlog.Error("create cluster err: %v", err.Error())
		response.FailReturn(c, errcode.CreateResourceError("cluster"))
		return
	}

	response.SuccessReturn(c, nil)
	return
}

func UpdateCluster(c *gin.Context, svc *service.ServiceContext) {

	name := c.Param("name")

	clusterRequest := model.Cluster{}
	err := c.BindJSON(&clusterRequest)
	if err != nil {
		mlog.Error("get create cluster body err:%s", err.Error())
		response.FailReturn(c, errcode.CreateResourceError("cluster"))
		return
	}

	cluster, err := svc.Model.FindOne(name)
	if err != nil || cluster == nil {
		mlog.Error("failed to get cluster %s err: %v", name, err.Error())
		response.FailReturn(c, errcode.ClusterNotFoundError(name))
		return
	}

	err = svc.Model.Update(clusterRequest)
	if err != nil {
		mlog.Error("update cluster err: %v", err.Error())
		response.FailReturn(c, errcode.UpdateResourceError("cluster"))
		return
	}

	response.SuccessReturn(c, nil)
	return
}

func DeleteCluster(c *gin.Context, svc *service.ServiceContext) {

	name := c.Param("name")

	cluster, err := svc.Model.FindOne(name)
	if err != nil || cluster == nil {
		mlog.Error("failed to get cluster %s err: %v", name, err.Error())
		response.FailReturn(c, errcode.ClusterNotFoundError(name))
		return
	}

	err = svc.Model.Delete(name)
	if err != nil {
		mlog.Error("update cluster err: %v", err.Error())
		response.FailReturn(c, errcode.UpdateResourceError("cluster"))
		return
	}

	response.SuccessReturn(c, nil)
	return
}
