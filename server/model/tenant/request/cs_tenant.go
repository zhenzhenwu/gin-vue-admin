package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tenant"
	"time"
)

type CsTenantSearch struct {
	tenant.CsTenant
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartEndTime   *time.Time `json:"startEndTime" form:"startEndTime"`
	EndEndTime     *time.Time `json:"endEndTime" form:"endEndTime"`
	request.PageInfo
}

type CsTenantApisReq struct {
	TenantId uint   `json:"tenantId" form:"tenantId"`
	ApiIDs   []uint `json:"apiIds" form:"apiIds"`
}

type CsTenantMenusReq struct {
	TenantId uint   `json:"tenantId" form:"tenantId" `
	MenuIDs  []uint `json:"menuIds" form:"menuIds"`
}
