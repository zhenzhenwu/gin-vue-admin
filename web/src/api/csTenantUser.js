import service from '@/utils/request'

// @Tags CsTenantUser
// @Summary 创建CsTenantUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsTenantUser true "创建CsTenantUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csTenantUser/createCsTenantUser [post]
export const createCsTenantUser = (data) => {
  return service({
    url: '/csTenantUser/createCsTenantUser',
    method: 'post',
    data
  })
}

// @Tags CsTenantUser
// @Summary 删除CsTenantUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsTenantUser true "删除CsTenantUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csTenantUser/deleteCsTenantUser [delete]
export const deleteCsTenantUser = (data) => {
  return service({
    url: '/csTenantUser/deleteCsTenantUser',
    method: 'delete',
    data
  })
}

// @Tags CsTenantUser
// @Summary 删除CsTenantUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CsTenantUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csTenantUser/deleteCsTenantUser [delete]
export const deleteCsTenantUserByIds = (data) => {
  return service({
    url: '/csTenantUser/deleteCsTenantUserByIds',
    method: 'delete',
    data
  })
}

// @Tags CsTenantUser
// @Summary 更新CsTenantUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsTenantUser true "更新CsTenantUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /csTenantUser/updateCsTenantUser [put]
export const updateCsTenantUser = (data) => {
  return service({
    url: '/csTenantUser/updateCsTenantUser',
    method: 'put',
    data
  })
}

// @Tags CsTenantUser
// @Summary 用id查询CsTenantUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CsTenantUser true "用id查询CsTenantUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /csTenantUser/findCsTenantUser [get]
export const findCsTenantUser = (params) => {
  return service({
    url: '/csTenantUser/findCsTenantUser',
    method: 'get',
    params
  })
}

// @Tags CsTenantUser
// @Summary 分页获取CsTenantUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CsTenantUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csTenantUser/getCsTenantUserList [get]
export const getCsTenantUserList = (params) => {
  return service({
    url: '/csTenantUser/getCsTenantUserList',
    method: 'get',
    params
  })
}
