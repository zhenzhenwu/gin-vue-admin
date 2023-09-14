import service from '@/utils/request'

// @Tags CsOperatorAuthorities
// @Summary 创建CsOperatorAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsOperatorAuthorities true "创建CsOperatorAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csOperatorAuthorities/createCsOperatorAuthorities [post]
export const createCsOperatorAuthorities = (data) => {
  return service({
    url: '/csOperatorAuthorities/createCsOperatorAuthorities',
    method: 'post',
    data
  })
}

// @Tags CsOperatorAuthorities
// @Summary 删除CsOperatorAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsOperatorAuthorities true "删除CsOperatorAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csOperatorAuthorities/deleteCsOperatorAuthorities [delete]
export const deleteCsOperatorAuthorities = (data) => {
  return service({
    url: '/csOperatorAuthorities/deleteCsOperatorAuthorities',
    method: 'delete',
    data
  })
}

// @Tags CsOperatorAuthorities
// @Summary 删除CsOperatorAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CsOperatorAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csOperatorAuthorities/deleteCsOperatorAuthorities [delete]
export const deleteCsOperatorAuthoritiesByIds = (data) => {
  return service({
    url: '/csOperatorAuthorities/deleteCsOperatorAuthoritiesByIds',
    method: 'delete',
    data
  })
}

// @Tags CsOperatorAuthorities
// @Summary 更新CsOperatorAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsOperatorAuthorities true "更新CsOperatorAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /csOperatorAuthorities/updateCsOperatorAuthorities [put]
export const updateCsOperatorAuthorities = (data) => {
  return service({
    url: '/csOperatorAuthorities/updateCsOperatorAuthorities',
    method: 'put',
    data
  })
}

// @Tags CsOperatorAuthorities
// @Summary 用id查询CsOperatorAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CsOperatorAuthorities true "用id查询CsOperatorAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /csOperatorAuthorities/findCsOperatorAuthorities [get]
export const findCsOperatorAuthorities = (params) => {
  return service({
    url: '/csOperatorAuthorities/findCsOperatorAuthorities',
    method: 'get',
    params
  })
}

// @Tags CsOperatorAuthorities
// @Summary 分页获取CsOperatorAuthorities列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CsOperatorAuthorities列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csOperatorAuthorities/getCsOperatorAuthoritiesList [get]
export const getCsOperatorAuthoritiesList = (params) => {
  return service({
    url: '/csOperatorAuthorities/getCsOperatorAuthoritiesList',
    method: 'get',
    params
  })
}
