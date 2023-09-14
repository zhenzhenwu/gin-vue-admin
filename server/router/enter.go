package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/operator"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/tenant"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Tenant   tenant.RouterGroup
	Operator operator.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
