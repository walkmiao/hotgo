package router

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/controller/business/car"
	"hotgo/internal/controller/business/customer"
	"hotgo/internal/service"
	"hotgo/utility/simple"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Business 业务路由
func Business(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group(simple.RouterPrefix(ctx, consts.Business), func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().BusinessAuth)
		group.Bind(
			customer.Customer,
			car.CarModel,
		)
	})
}
