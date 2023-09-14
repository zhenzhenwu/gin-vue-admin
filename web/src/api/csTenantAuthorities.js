import service from '@/utils/request'

// @Tags CsTenantAuthorities
// @Summary 创建CsTenantAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsTenantAuthorities true "创建CsTenantAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csTenantAuthorities/createCsTenantAuthorities [post]
export const createCsTenantAuthorities = (data) => {
  return service({
    url: '/csTenantAuthorities/createCsTenantAuthorities',
    method: 'post',
    data
  })
}

// @Tags CsTenantAuthorities
// @Summary 删除CsTenantAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsTenantAuthorities true "删除CsTenantAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csTenantAuthorities/deleteCsTenantAuthorities [delete]
export const deleteCsTenantAuthorities = (data) => {
  return service({
    url: '/csTenantAuthorities/deleteCsTenantAuthorities',
    method: 'delete',
    data
  })
}

// @Tags CsTenantAuthorities
// @Summary 删除CsTenantAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CsTenantAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csTenantAuthorities/deleteCsTenantAuthorities [delete]
export const deleteCsTenantAuthoritiesByIds = (data) => {
  return service({
    url: '/csTenantAuthorities/deleteCsTenantAuthoritiesByIds',
    method: 'delete',
    data
  })
}

// @Tags CsTenantAuthorities
// @Summary 更新CsTenantAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsTenantAuthorities true "更新CsTenantAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /csTenantAuthorities/updateCsTenantAuthorities [put]
export const updateCsTenantAuthorities = (data) => {
  return service({
    url: '/csTenantAuthorities/updateCsTenantAuthorities',
    method: 'put',
    data
  })
}

// @Tags CsTenantAuthorities
// @Summary 用id查询CsTenantAuthorities
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CsTenantAuthorities true "用id查询CsTenantAuthorities"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /csTenantAuthorities/findCsTenantAuthorities [get]
export const findCsTenantAuthorities = (params) => {
  return service({
    url: '/csTenantAuthorities/findCsTenantAuthorities',
    method: 'get',
    params
  })
}

// @Tags CsTenantAuthorities
// @Summary 分页获取CsTenantAuthorities列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CsTenantAuthorities列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csTenantAuthorities/getCsTenantAuthoritiesList [get]
export const getCsTenantAuthoritiesList = (params) => {
  return service({
    url: '/csTenantAuthorities/getCsTenantAuthoritiesList',
    method: 'get',
    params
  })
}
