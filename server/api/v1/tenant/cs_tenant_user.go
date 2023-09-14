package tenant

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tenant"
	tenantReq "github.com/flipped-aurora/gin-vue-admin/server/model/tenant/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CsTenantUserApi struct {
}

var csTenantUserService = service.ServiceGroupApp.TenantServiceGroup.CsTenantUserService

// CreateCsTenantUser 创建CsTenantUser
// @Tags CsTenantUser
// @Summary 创建CsTenantUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tenant.CsTenantUser true "创建CsTenantUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csTenantUser/createCsTenantUser [post]
func (csTenantUserApi *CsTenantUserApi) CreateCsTenantUser(c *gin.Context) {
	var csTenantUser tenant.CsTenantUser
	err := c.ShouldBindJSON(&csTenantUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	csTenantUser.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Username":          {utils.NotEmpty()},
		"Password":          {utils.NotEmpty()},
		"NickName":          {utils.NotEmpty()},
		"TenantAuthorityId": {utils.NotEmpty()},
	}
	if err := utils.Verify(csTenantUser, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csTenantUserService.CreateCsTenantUser(&csTenantUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCsTenantUser 删除CsTenantUser
// @Tags CsTenantUser
// @Summary 删除CsTenantUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tenant.CsTenantUser true "删除CsTenantUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csTenantUser/deleteCsTenantUser [delete]
func (csTenantUserApi *CsTenantUserApi) DeleteCsTenantUser(c *gin.Context) {
	var csTenantUser tenant.CsTenantUser
	err := c.ShouldBindJSON(&csTenantUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	csTenantUser.DeletedBy = utils.GetUserID(c)
	if err := csTenantUserService.DeleteCsTenantUser(csTenantUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCsTenantUserByIds 批量删除CsTenantUser
// @Tags CsTenantUser
// @Summary 批量删除CsTenantUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CsTenantUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /csTenantUser/deleteCsTenantUserByIds [delete]
func (csTenantUserApi *CsTenantUserApi) DeleteCsTenantUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := csTenantUserService.DeleteCsTenantUserByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCsTenantUser 更新CsTenantUser
// @Tags CsTenantUser
// @Summary 更新CsTenantUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tenant.CsTenantUser true "更新CsTenantUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /csTenantUser/updateCsTenantUser [put]
func (csTenantUserApi *CsTenantUserApi) UpdateCsTenantUser(c *gin.Context) {
	var csTenantUser tenant.CsTenantUser
	err := c.ShouldBindJSON(&csTenantUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	csTenantUser.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Username":          {utils.NotEmpty()},
		"Password":          {utils.NotEmpty()},
		"NickName":          {utils.NotEmpty()},
		"TenantAuthorityId": {utils.NotEmpty()},
	}
	if err := utils.Verify(csTenantUser, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csTenantUserService.UpdateCsTenantUser(csTenantUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCsTenantUser 用id查询CsTenantUser
// @Tags CsTenantUser
// @Summary 用id查询CsTenantUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tenant.CsTenantUser true "用id查询CsTenantUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /csTenantUser/findCsTenantUser [get]
func (csTenantUserApi *CsTenantUserApi) FindCsTenantUser(c *gin.Context) {
	var csTenantUser tenant.CsTenantUser
	err := c.ShouldBindQuery(&csTenantUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recsTenantUser, err := csTenantUserService.GetCsTenantUser(csTenantUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recsTenantUser": recsTenantUser}, c)
	}
}

// GetCsTenantUserList 分页获取CsTenantUser列表
// @Tags CsTenantUser
// @Summary 分页获取CsTenantUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tenantReq.CsTenantUserSearch true "分页获取CsTenantUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csTenantUser/getCsTenantUserList [get]
func (csTenantUserApi *CsTenantUserApi) GetCsTenantUserList(c *gin.Context) {
	var pageInfo tenantReq.CsTenantUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := csTenantUserService.GetCsTenantUserInfoList(pageInfo); err != nil {
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
