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

type CsTenantAuthoritiesApi struct {
}

var csTenantAuthoritiesService = service.ServiceGroupApp.TenantServiceGroup.CsTenantAuthoritiesService

// CreateCsTenantAuthorities 创建CsTenantAuthorities
// @Tags CsTenantAuthorities
// @Summary 创建CsTenantAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tenant.CsTenantAuthorities true "创建CsTenantAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csTenantAuthorities/createCsTenantAuthorities [post]
func (csTenantAuthoritiesApi *CsTenantAuthoritiesApi) CreateCsTenantAuthorities(c *gin.Context) {
	var csTenantAuthorities tenant.CsTenantAuthorities
	err := c.ShouldBindJSON(&csTenantAuthorities)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	csTenantAuthorities.CreatedBy = utils.GetUserID(c)
	if err := csTenantAuthoritiesService.CreateCsTenantAuthorities(&csTenantAuthorities); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCsTenantAuthorities 删除CsTenantAuthorities
// @Tags CsTenantAuthorities
// @Summary 删除CsTenantAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tenant.CsTenantAuthorities true "删除CsTenantAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csTenantAuthorities/deleteCsTenantAuthorities [delete]
func (csTenantAuthoritiesApi *CsTenantAuthoritiesApi) DeleteCsTenantAuthorities(c *gin.Context) {
	var csTenantAuthorities tenant.CsTenantAuthorities
	err := c.ShouldBindJSON(&csTenantAuthorities)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	csTenantAuthorities.DeletedBy = utils.GetUserID(c)
	if err := csTenantAuthoritiesService.DeleteCsTenantAuthorities(csTenantAuthorities); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCsTenantAuthoritiesByIds 批量删除CsTenantAuthorities
// @Tags CsTenantAuthorities
// @Summary 批量删除CsTenantAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CsTenantAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /csTenantAuthorities/deleteCsTenantAuthoritiesByIds [delete]
func (csTenantAuthoritiesApi *CsTenantAuthoritiesApi) DeleteCsTenantAuthoritiesByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := csTenantAuthoritiesService.DeleteCsTenantAuthoritiesByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCsTenantAuthorities 更新CsTenantAuthorities
// @Tags CsTenantAuthorities
// @Summary 更新CsTenantAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tenant.CsTenantAuthorities true "更新CsTenantAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /csTenantAuthorities/updateCsTenantAuthorities [put]
func (csTenantAuthoritiesApi *CsTenantAuthoritiesApi) UpdateCsTenantAuthorities(c *gin.Context) {
	var csTenantAuthorities tenant.CsTenantAuthorities
	err := c.ShouldBindJSON(&csTenantAuthorities)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	csTenantAuthorities.UpdatedBy = utils.GetUserID(c)
	if err := csTenantAuthoritiesService.UpdateCsTenantAuthorities(csTenantAuthorities); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCsTenantAuthorities 用id查询CsTenantAuthorities
// @Tags CsTenantAuthorities
// @Summary 用id查询CsTenantAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tenant.CsTenantAuthorities true "用id查询CsTenantAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /csTenantAuthorities/findCsTenantAuthorities [get]
func (csTenantAuthoritiesApi *CsTenantAuthoritiesApi) FindCsTenantAuthorities(c *gin.Context) {
	var csTenantAuthorities tenant.CsTenantAuthorities
	err := c.ShouldBindQuery(&csTenantAuthorities)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recsTenantAuthorities, err := csTenantAuthoritiesService.GetCsTenantAuthorities(csTenantAuthorities.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recsTenantAuthorities": recsTenantAuthorities}, c)
	}
}

// GetCsTenantAuthoritiesList 分页获取CsTenantAuthorities列表
// @Tags CsTenantAuthorities
// @Summary 分页获取CsTenantAuthorities列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tenantReq.CsTenantAuthoritiesSearch true "分页获取CsTenantAuthorities列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csTenantAuthorities/getCsTenantAuthoritiesList [get]
func (csTenantAuthoritiesApi *CsTenantAuthoritiesApi) GetCsTenantAuthoritiesList(c *gin.Context) {
	var pageInfo tenantReq.CsTenantAuthoritiesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := csTenantAuthoritiesService.GetCsTenantAuthoritiesInfoList(pageInfo); err != nil {
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
