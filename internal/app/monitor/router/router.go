package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/yizhibian/gdumper/internal/app/monitor/controller"
	"github.com/yizhibian/gdumper/internal/app/system/service"
	"github.com/yizhibian/gdumper/library/libRouter"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/kube", func(group *ghttp.RouterGroup) {
		//登录验证拦截
		service.GfToken().Middleware(group)
		//context拦截器
		group.Middleware(service.Middleware().Ctx, service.Middleware().Auth)
		//后台操作日志记录
		group.Hook("/*", ghttp.HookAfterOutput, service.OperateLog().OperationLog)
		group.Bind(
			controller.Kube,
		)
		//自动绑定定义的控制器
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
