package tenant

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CsTenantAuthoritiesRouter struct {
}

// InitCsTenantAuthoritiesRouter 初始化 CsTenantAuthorities 路由信息
func (s *CsTenantAuthoritiesRouter) InitCsTenantAuthoritiesRouter(Router *gin.RouterGroup) {
	csTenantAuthoritiesRouter := Router.Group("csTenantAuthorities").Use(middleware.OperationRecord())
	csTenantAuthoritiesRouterWithoutRecord := Router.Group("csTenantAuthorities")
	var csTenantAuthoritiesApi = v1.ApiGroupApp.TenantApiGroup.CsTenantAuthoritiesApi
	{
		csTenantAuthoritiesRouter.POST("createCsTenantAuthorities", csTenantAuthoritiesApi.CreateCsTenantAuthorities)             // 新建CsTenantAuthorities
		csTenantAuthoritiesRouter.DELETE("deleteCsTenantAuthorities", csTenantAuthoritiesApi.DeleteCsTenantAuthorities)           // 删除CsTenantAuthorities
		csTenantAuthoritiesRouter.DELETE("deleteCsTenantAuthoritiesByIds", csTenantAuthoritiesApi.DeleteCsTenantAuthoritiesByIds) // 批量删除CsTenantAuthorities
		csTenantAuthoritiesRouter.PUT("updateCsTenantAuthorities", csTenantAuthoritiesApi.UpdateCsTenantAuthorities)              // 更新CsTenantAuthorities
	}
	{
		csTenantAuthoritiesRouterWithoutRecord.GET("findCsTenantAuthorities", csTenantAuthoritiesApi.FindCsTenantAuthorities)       // 根据ID获取CsTenantAuthorities
		csTenantAuthoritiesRouterWithoutRecord.GET("getCsTenantAuthoritiesList", csTenantAuthoritiesApi.GetCsTenantAuthoritiesList) // 获取CsTenantAuthorities列表
	}
}
