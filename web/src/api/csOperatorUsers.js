import service from '@/utils/request'

// @Tags CsOperatorUsers
// @Summary 创建CsOperatorUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsOperatorUsers true "创建CsOperatorUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csOperatorUsers/createCsOperatorUsers [post]
export const createCsOperatorUsers = (data) => {
  return service({
    url: '/csOperatorUsers/createCsOperatorUsers',
    method: 'post',
    data
  })
}

// @Tags CsOperatorUsers
// @Summary 删除CsOperatorUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsOperatorUsers true "删除CsOperatorUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csOperatorUsers/deleteCsOperatorUsers [delete]
export const deleteCsOperatorUsers = (data) => {
  return service({
    url: '/csOperatorUsers/deleteCsOperatorUsers',
    method: 'delete',
    data
  })
}

// @Tags CsOperatorUsers
// @Summary 删除CsOperatorUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CsOperatorUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /csOperatorUsers/deleteCsOperatorUsers [delete]
export const deleteCsOperatorUsersByIds = (data) => {
  return service({
    url: '/csOperatorUsers/deleteCsOperatorUsersByIds',
    method: 'delete',
    data
  })
}

// @Tags CsOperatorUsers
// @Summary 更新CsOperatorUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CsOperatorUsers true "更新CsOperatorUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /csOperatorUsers/updateCsOperatorUsers [put]
export const updateCsOperatorUsers = (data) => {
  return service({
    url: '/csOperatorUsers/updateCsOperatorUsers',
    method: 'put',
    data
  })
}

// @Tags CsOperatorUsers
// @Summary 用id查询CsOperatorUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CsOperatorUsers true "用id查询CsOperatorUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /csOperatorUsers/findCsOperatorUsers [get]
export const findCsOperatorUsers = (params) => {
  return service({
    url: '/csOperatorUsers/findCsOperatorUsers',
    method: 'get',
    params
  })
}

// @Tags CsOperatorUsers
// @Summary 分页获取CsOperatorUsers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CsOperatorUsers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /csOperatorUsers/getCsOperatorUsersList [get]
export const getCsOperatorUsersList = (params) => {
  return service({
    url: '/csOperatorUsers/getCsOperatorUsersList',
    method: 'get',
    params
  })
}
