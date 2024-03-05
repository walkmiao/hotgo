// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package middleware

import (
	"hotgo/internal/consts"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/response"
	"hotgo/internal/service"
	"hotgo/utility/simple"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// ApiAuth API鉴权中间件
func (s *sMiddleware) ApiAuth(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		path = gstr.Replace(r.URL.Path, simple.RouterPrefix(ctx, consts.AppApi), "", 1)
	)

	// 不需要验证登录的路由地址
	if s.IsExceptLogin(ctx, consts.AppApi, path) {
		r.Middleware.Next()
		return
	}

	// 将用户信息传递到上下文中
	if err := s.DeliverUserContext(r); err != nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
		return
	}

	// 验证路由访问权限
	// ...

	r.Middleware.Next()
}

func (s *sMiddleware) BusinessAuth(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		path = gstr.Replace(r.URL.Path, simple.RouterPrefix(ctx, consts.Business), "", 1)
	)

	// 不需要验证登录的路由地址
	if s.IsExceptLogin(ctx, consts.Business, path) {
		r.Middleware.Next()
		return
	}

	// 将用户信息传递到上下文中
	if err := s.DeliverUserContext(r); err != nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
		return
	}

	// 验证路由访问权限
	if !service.AdminRole().Verify(ctx, path, r.Method) {
		g.Log().Debugf(ctx, "AdminAuth fail path:%+v, GetRoleKey:%+v, r.Method:%+v", path, contexts.GetRoleKey(ctx), r.Method)
		response.JsonExit(r, gcode.CodeSecurityReason.Code(), "你没有访问权限！")
		return
	}
	// 验证路由访问权限
	// ...

	r.Middleware.Next()
}
