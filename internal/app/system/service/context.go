// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/yizhibian/gdumper/internal/app/system/model"
)

type IContext interface {
	Init(r *ghttp.Request, customCtx *model.Context)
	Get(ctx context.Context) *model.Context
	SetUser(ctx context.Context, ctxUser *model.ContextUser)
	GetLoginUser(ctx context.Context) *model.ContextUser
	GetUserId(ctx context.Context) uint64
}

var localContext IContext

func Context() IContext {
	if localContext == nil {
		panic("implement not found for interface IContext, forgot register?")
	}
	return localContext
}

func RegisterContext(i IContext) {
	localContext = i
}
