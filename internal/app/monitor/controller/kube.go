package controller

import (
	"context"
	"github.com/yizhibian/gdumper/api/v1/monitor"
	"github.com/yizhibian/gdumper/internal/app/monitor/service"
)

var Kube = kubeController{}

type kubeController struct {
	BaseController
}

// List 部门列表
func (c *kubeController) List(ctx context.Context, req *monitor.PodListReq) (res *monitor.PodListRes, err error) {
	res = new(monitor.PodListRes)
	var namespace string
	if len(req.Namespace) == 0 {
		namespace = "default"
	} else {
		namespace = req.Namespace
	}
	res, err = service.Kube().GetPodsByNamespace(ctx, namespace)
	return
}

// GetShit 测试
func (c *kubeController) GetShit(ctx context.Context, req *monitor.GetShitReq) (res *monitor.GetShitRes, err error) {
	err = service.Kube().GetSomeShit(ctx, "shit")
	return
}
