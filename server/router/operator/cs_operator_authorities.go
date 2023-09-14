package operator

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CsOperatorAuthoritiesRouter struct {
}

// InitCsOperatorAuthoritiesRouter 初始化 CsOperatorAuthorities 路由信息
func (s *CsOperatorAuthoritiesRouter) InitCsOperatorAuthoritiesRouter(Router *gin.RouterGroup) {
	csOperatorAuthoritiesRouter := Router.Group("csOperatorAuthorities").Use(middleware.OperationRecord())
	csOperatorAuthoritiesRouterWithoutRecord := Router.Group("csOperatorAuthorities")
	var csOperatorAuthoritiesApi = v1.ApiGroupApp.OperatorApiGroup.CsOperatorAuthoritiesApi
	{
		csOperatorAuthoritiesRouter.POST("createCsOperatorAuthorities", csOperatorAuthoritiesApi.CreateCsOperatorAuthorities)             // 新建CsOperatorAuthorities
		csOperatorAuthoritiesRouter.DELETE("deleteCsOperatorAuthorities", csOperatorAuthoritiesApi.DeleteCsOperatorAuthorities)           // 删除CsOperatorAuthorities
		csOperatorAuthoritiesRouter.DELETE("deleteCsOperatorAuthoritiesByIds", csOperatorAuthoritiesApi.DeleteCsOperatorAuthoritiesByIds) // 批量删除CsOperatorAuthorities
		csOperatorAuthoritiesRouter.PUT("updateCsOperatorAuthorities", csOperatorAuthoritiesApi.UpdateCsOperatorAuthorities)              // 更新CsOperatorAuthorities
	}
	{
		csOperatorAuthoritiesRouterWithoutRecord.GET("findCsOperatorAuthorities", csOperatorAuthoritiesApi.FindCsOperatorAuthorities)       // 根据ID获取CsOperatorAuthorities
		csOperatorAuthoritiesRouterWithoutRecord.GET("getCsOperatorAuthoritiesList", csOperatorAuthoritiesApi.GetCsOperatorAuthoritiesList) // 获取CsOperatorAuthorities列表
	}
}
