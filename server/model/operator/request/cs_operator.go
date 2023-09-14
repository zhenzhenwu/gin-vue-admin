package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operator"
	"time"
)

type CsOperatorSearch struct {
	operator.CsOperator
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type CsOperatorApisReq struct {
	OperatorId uint   `json:"operatorId" form:"operatorId"`
	ApiIDs     []uint `json:"apiIds" form:"apiIds"`
}

type CsOperatorMenusReq struct {
	OperatorId uint   `json:"operatorId" form:"operatorId" `
	MenuIDs    []uint `json:"menuIds" form:"menuIds"`
}
