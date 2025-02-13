package monitor

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/yizhibian/gdumper/api/v1/common"
	"github.com/yizhibian/gdumper/internal/app/monitor/model"
)

type PodListReq struct {
	g.Meta    `path:"/pods/list" tags:"集群管理" method:"get" summary:"集群列表"`
	Namespace string `json:"namespace"`
	commonApi.PageReq
}

type PodListRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.PodInfo `json:"list"`
}

type GetShitReq struct {
	g.Meta `path:"/pods/shit" tags:"集群管理" method:"get" summary:"集群列表"`
	commonApi.PageReq
}

type GetShitRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
}
