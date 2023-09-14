package operator

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operator"
	operatorReq "github.com/flipped-aurora/gin-vue-admin/server/model/operator/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CsOperatorUsersApi struct {
}

var csOperatorUsersService = service.ServiceGroupApp.OperatorServiceGroup.CsOperatorUsersService

// CreateCsOperatorUsers 创建CsOperatorUsers
// @Tags CsOperatorUsers
// @Summary 创建CsOperatorUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body operator.CsOperatorUsers true "创建CsOperatorUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csOperatorUsers/createCsOperatorUsers [post]
func (csOperatorUsersApi *CsOperatorUsersApi) CreateCsOperatorUsers(c *gin.Context) {
	var csOperatorUsers operator.CsOperatorUsers
	err := c.ShouldBindJSON(&csOperatorUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	csOperatorUsers.CreatedBy = utils.GetUserID(c)
	if err := csOperatorUsersService.CreateCsOperatorUsers(&csOperatorUsers); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCsOperatorUsers 删除CsOperatorUsers
// @Tags CsOperatorUsers
// @Summary 删除CsOperatorUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body operator.CsOperatorUsers true "删除CsOperatorUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csOperatorUsers/deleteCsOperatorUsers [delete]
func (csOperatorUsersApi *CsOperatorUsersApi) DeleteCsOperatorUsers(c *gin.Context) {
	var csOperatorUsers operator.CsOperatorUsers
	err := c.ShouldBindJSON(&csOperatorUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	csOperatorUsers.DeletedBy = utils.GetUserID(c)
	if err := csOperatorUsersService.DeleteCsOperatorUsers(csOperatorUsers); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCsOperatorUsersByIds 批量删除CsOperatorUsers
// @Tags CsOperatorUsers
// @Summary 批量删除CsOperatorUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CsOperatorUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /csOperatorUsers/deleteCsOperatorUsersByIds [delete]
func (csOperatorUsersApi *CsOperatorUsersApi) DeleteCsOperatorUsersByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := csOperatorUsersService.DeleteCsOperatorUsersByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCsOperatorUsers 更新CsOperatorUsers
// @Tags CsOperatorUsers
// @Summary 更新CsOperatorUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body operator.CsOperatorUsers true "更新CsOperatorUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /csOperatorUsers/updateCsOperatorUsers [put]
func (csOperatorUsersApi *CsOperatorUsersApi) UpdateCsOperatorUsers(c *gin.Context) {
	var csOperatorUsers operator.CsOperatorUsers
	err := c.ShouldBindJSON(&csOperatorUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	csOperatorUsers.UpdatedBy = utils.GetUserID(c)
	if err := csOperatorUsersService.UpdateCsOperatorUsers(csOperatorUsers); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCsOperatorUsers 用id查询CsOperatorUsers
// @Tags CsOperatorUsers
// @Summary 用id查询CsOperatorUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query operator.CsOperatorUsers true "用id查询CsOperatorUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /csOperatorUsers/findCsOperatorUsers [get]
func (csOperatorUsersApi *CsOperatorUsersApi) FindCsOperatorUsers(c *gin.Context) {
	var csOperatorUsers operator.CsOperatorUsers
	err := c.ShouldBindQuery(&csOperatorUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recsOperatorUsers, err := csOperatorUsersService.GetCsOperatorUsers(csOperatorUsers.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recsOperatorUsers": recsOperatorUsers}, c)
	}
}

// GetCsOperatorUsersList 分页获取CsOperatorUsers列表
// @Tags CsOperatorUsers
// @Summary 分页获取CsOperatorUsers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query operatorReq.CsOperatorUsersSearch true "分页获取CsOperatorUsers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csOperatorUsers/getCsOperatorUsersList [get]
func (csOperatorUsersApi *CsOperatorUsersApi) GetCsOperatorUsersList(c *gin.Context) {
	var pageInfo operatorReq.CsOperatorUsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := csOperatorUsersService.GetCsOperatorUsersInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
