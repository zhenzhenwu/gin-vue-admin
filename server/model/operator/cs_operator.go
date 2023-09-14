// 自动生成模板CsOperator
package operator

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// CsOperator 结构体
type CsOperator struct {
	global.GVA_MODEL
	Username             string     `json:"username" form:"username" gorm:"column:username;comment:租户名;size:191;"`
	Password             string     `json:"password" form:"password" gorm:"column:password;comment:密码;size:191;"`
	Nickname             string     `json:"nickname" form:"nickname" gorm:"column:nickname;comment:租户昵称;size:191;"`
	EndTime              *time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:;"`
	Allocation           *int       `json:"allocation" form:"allocation" gorm:"column:allocation;comment:;size:19;"`
	AllocationProportion *float64   `json:"allocationProportion" form:"allocationProportion" gorm:"column:allocation_proportion;comment:;size:22;"`
	CreatedBy            uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy            uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy            uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName CsOperator 表名
func (CsOperator) TableName() string {
	return "cs_operator"
}
