package operator

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CsOperatorUsersRouter struct {
}

// InitCsOperatorUsersRouter 初始化 CsOperatorUsers 路由信息
func (s *CsOperatorUsersRouter) InitCsOperatorUsersRouter(Router *gin.RouterGroup) {
	csOperatorUsersRouter := Router.Group("csOperatorUsers").Use(middleware.OperationRecord())
	csOperatorUsersRouterWithoutRecord := Router.Group("csOperatorUsers")
	var csOperatorUsersApi = v1.ApiGroupApp.OperatorApiGroup.CsOperatorUsersApi
	{
		csOperatorUsersRouter.POST("createCsOperatorUsers", csOperatorUsersApi.CreateCsOperatorUsers)             // 新建CsOperatorUsers
		csOperatorUsersRouter.DELETE("deleteCsOperatorUsers", csOperatorUsersApi.DeleteCsOperatorUsers)           // 删除CsOperatorUsers
		csOperatorUsersRouter.DELETE("deleteCsOperatorUsersByIds", csOperatorUsersApi.DeleteCsOperatorUsersByIds) // 批量删除CsOperatorUsers
		csOperatorUsersRouter.PUT("updateCsOperatorUsers", csOperatorUsersApi.UpdateCsOperatorUsers)              // 更新CsOperatorUsers
	}
	{
		csOperatorUsersRouterWithoutRecord.GET("findCsOperatorUsers", csOperatorUsersApi.FindCsOperatorUsers)       // 根据ID获取CsOperatorUsers
		csOperatorUsersRouterWithoutRecord.GET("getCsOperatorUsersList", csOperatorUsersApi.GetCsOperatorUsersList) // 获取CsOperatorUsers列表
	}
}
