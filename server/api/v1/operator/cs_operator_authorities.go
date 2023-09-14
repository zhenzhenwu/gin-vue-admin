package operator

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operator"
	operatorReq "github.com/flipped-aurora/gin-vue-admin/server/model/operator/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CsOperatorAuthoritiesApi struct {
}

var csOperatorAuthoritiesService = service.ServiceGroupApp.OperatorServiceGroup.CsOperatorAuthoritiesService

// CreateCsOperatorAuthorities 创建CsOperatorAuthorities
// @Tags CsOperatorAuthorities
// @Summary 创建CsOperatorAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body operator.CsOperatorAuthorities true "创建CsOperatorAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csOperatorAuthorities/createCsOperatorAuthorities [post]
func (csOperatorAuthoritiesApi *CsOperatorAuthoritiesApi) CreateCsOperatorAuthorities(c *gin.Context) {
	var csOperatorAuthorities operator.CsOperatorAuthorities
	err := c.ShouldBindJSON(&csOperatorAuthorities)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csOperatorAuthoritiesService.CreateCsOperatorAuthorities(&csOperatorAuthorities); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCsOperatorAuthorities 删除CsOperatorAuthorities
// @Tags CsOperatorAuthorities
// @Summary 删除CsOperatorAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body operator.CsOperatorAuthorities true "删除CsOperatorAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csOperatorAuthorities/deleteCsOperatorAuthorities [delete]
func (csOperatorAuthoritiesApi *CsOperatorAuthoritiesApi) DeleteCsOperatorAuthorities(c *gin.Context) {
	var csOperatorAuthorities operator.CsOperatorAuthorities
	err := c.ShouldBindJSON(&csOperatorAuthorities)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csOperatorAuthoritiesService.DeleteCsOperatorAuthorities(csOperatorAuthorities); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCsOperatorAuthoritiesByIds 批量删除CsOperatorAuthorities
// @Tags CsOperatorAuthorities
// @Summary 批量删除CsOperatorAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CsOperatorAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /csOperatorAuthorities/deleteCsOperatorAuthoritiesByIds [delete]
func (csOperatorAuthoritiesApi *CsOperatorAuthoritiesApi) DeleteCsOperatorAuthoritiesByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csOperatorAuthoritiesService.DeleteCsOperatorAuthoritiesByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCsOperatorAuthorities 更新CsOperatorAuthorities
// @Tags CsOperatorAuthorities
// @Summary 更新CsOperatorAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body operator.CsOperatorAuthorities true "更新CsOperatorAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /csOperatorAuthorities/updateCsOperatorAuthorities [put]
func (csOperatorAuthoritiesApi *CsOperatorAuthoritiesApi) UpdateCsOperatorAuthorities(c *gin.Context) {
	var csOperatorAuthorities operator.CsOperatorAuthorities
	err := c.ShouldBindJSON(&csOperatorAuthorities)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csOperatorAuthoritiesService.UpdateCsOperatorAuthorities(csOperatorAuthorities); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCsOperatorAuthorities 用id查询CsOperatorAuthorities
// @Tags CsOperatorAuthorities
// @Summary 用id查询CsOperatorAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query operator.CsOperatorAuthorities true "用id查询CsOperatorAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /csOperatorAuthorities/findCsOperatorAuthorities [get]
func (csOperatorAuthoritiesApi *CsOperatorAuthoritiesApi) FindCsOperatorAuthorities(c *gin.Context) {
	var csOperatorAuthorities operator.CsOperatorAuthorities
	err := c.ShouldBindQuery(&csOperatorAuthorities)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recsOperatorAuthorities, err := csOperatorAuthoritiesService.GetCsOperatorAuthorities(csOperatorAuthorities.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recsOperatorAuthorities": recsOperatorAuthorities}, c)
	}
}

// GetCsOperatorAuthoritiesList 分页获取CsOperatorAuthorities列表
// @Tags CsOperatorAuthorities
// @Summary 分页获取CsOperatorAuthorities列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query operatorReq.CsOperatorAuthoritiesSearch true "分页获取CsOperatorAuthorities列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csOperatorAuthorities/getCsOperatorAuthoritiesList [get]
func (csOperatorAuthoritiesApi *CsOperatorAuthoritiesApi) GetCsOperatorAuthoritiesList(c *gin.Context) {
	var pageInfo operatorReq.CsOperatorAuthoritiesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := csOperatorAuthoritiesService.GetCsOperatorAuthoritiesInfoList(pageInfo); err != nil {
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
