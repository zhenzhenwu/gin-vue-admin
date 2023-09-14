// 自动生成模板CsOperatorAuthorities
package operator

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CsOperatorAuthorities 结构体
type CsOperatorAuthorities struct {
	global.GVA_MODEL
	AuthorityId   *int   `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:角色ID;size:20;"`
	AuthorityName string `json:"authorityName" form:"authorityName" gorm:"column:authority_name;comment:角色名;size:191;"`
	ParentId      *int   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父角色ID;size:20;"`
	DefaultRouter string `json:"defaultRouter" form:"defaultRouter" gorm:"column:default_router;comment:默认菜单;size:191;"`
}

// TableName CsOperatorAuthorities 表名
func (CsOperatorAuthorities) TableName() string {
	return "cs_operator_authorities"
}
