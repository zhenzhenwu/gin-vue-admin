package operator

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CsOperatorRouter struct {
}

// InitCsOperatorRouter 初始化 CsOperator 路由信息
func (s *CsOperatorRouter) InitCsOperatorRouter(Router *gin.RouterGroup) {
	csOperatorRouter := Router.Group("csOperator").Use(middleware.OperationRecord())
	csOperatorRouterWithoutRecord := Router.Group("csOperator")
	var csOperatorApi = v1.ApiGroupApp.OperatorApiGroup.CsOperatorApi
	{
		csOperatorRouter.POST("createCsOperator", csOperatorApi.CreateCsOperator)             // 新建CsOperator
		csOperatorRouter.DELETE("deleteCsOperator", csOperatorApi.DeleteCsOperator)           // 删除CsOperator
		csOperatorRouter.DELETE("deleteCsOperatorByIds", csOperatorApi.DeleteCsOperatorByIds) // 批量删除CsOperator
		csOperatorRouter.PUT("updateCsOperator", csOperatorApi.UpdateCsOperator)              // 更新CsOperator
		csOperatorRouter.PUT("setTenantApis", csOperatorApi.SetOperatorApis)                  // 系统分配运营商api
		csOperatorRouter.PUT("setTenantMenus", csOperatorApi.SetOperatorMenus)                // 系统分配运营商菜单
	}
	{
		csOperatorRouterWithoutRecord.GET("findCsOperator", csOperatorApi.FindCsOperator)       // 根据ID获取CsOperator
		csOperatorRouterWithoutRecord.GET("getCsOperatorList", csOperatorApi.GetCsOperatorList) // 获取CsOperator列表
	}
}
