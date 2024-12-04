package routers

import (
	"github.com/dinhdc/go-ecommerce/internal/routers/admin"
	"github.com/dinhdc/go-ecommerce/internal/routers/user"
)

type RouterGroup struct {
	User  user.UserRouterGroup
	Admin admin.AdminRouterGroup
}

var RouterGroupApi = new(RouterGroup)
