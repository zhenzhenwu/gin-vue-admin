package tenant

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CsTenantUserRouter struct {
}

// InitCsTenantUserRouter 初始化 CsTenantUser 路由信息
func (s *CsTenantUserRouter) InitCsTenantUserRouter(Router *gin.RouterGroup) {
	csTenantUserRouter := Router.Group("csTenantUser").Use(middleware.OperationRecord())
	csTenantUserRouterWithoutRecord := Router.Group("csTenantUser")
	var csTenantUserApi = v1.ApiGroupApp.TenantApiGroup.CsTenantUserApi
	{
		csTenantUserRouter.POST("createCsTenantUser", csTenantUserApi.CreateCsTenantUser)             // 新建CsTenantUser
		csTenantUserRouter.DELETE("deleteCsTenantUser", csTenantUserApi.DeleteCsTenantUser)           // 删除CsTenantUser
		csTenantUserRouter.DELETE("deleteCsTenantUserByIds", csTenantUserApi.DeleteCsTenantUserByIds) // 批量删除CsTenantUser
		csTenantUserRouter.PUT("updateCsTenantUser", csTenantUserApi.UpdateCsTenantUser)              // 更新CsTenantUser
	}
	{
		csTenantUserRouterWithoutRecord.GET("findCsTenantUser", csTenantUserApi.FindCsTenantUser)       // 根据ID获取CsTenantUser
		csTenantUserRouterWithoutRecord.GET("getCsTenantUserList", csTenantUserApi.GetCsTenantUserList) // 获取CsTenantUser列表
	}
}
