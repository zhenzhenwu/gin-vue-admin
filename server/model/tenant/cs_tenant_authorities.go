// 自动生成模板CsTenantAuthorities
package tenant

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CsTenantAuthorities 结构体
type CsTenantAuthorities struct {
	global.GVA_MODEL
	AuthorityId   *int   `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:角色ID;size:20;"`
	AuthorityName string `json:"authorityName" form:"authorityName" gorm:"column:authority_name;comment:角色名;size:191;"`
	DefaultRouter string `json:"defaultRouter" form:"defaultRouter" gorm:"column:default_router;comment:默认菜单;size:191;"`
	ParentId      *int   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父角色ID;size:20;"`
	CreatedBy     uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy     uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy     uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName CsTenantAuthorities 表名
func (CsTenantAuthorities) TableName() string {
	return "cs_tenant_authorities"
}
