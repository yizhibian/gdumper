// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/yizhibian/gdumper/api/v1/monitor"
)

type (
	IKube interface {
		GetSomeShit(c context.Context, namespace string) (err error)
		GetPodsByNamespace(c context.Context, namespace string) (res *monitor.PodListRes, err error)
	}
)

var (
	localKube IKube
)

func Kube() IKube {
	if localKube == nil {
		panic("implement not found for interface IKube, forgot register?")
	}
	return localKube
}

func RegisterKube(i IKube) {
	localKube = i
}
