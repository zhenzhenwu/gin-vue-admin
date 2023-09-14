// 自动生成模板CsOperatorUsers
package operator

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CsOperatorUsers 结构体
type CsOperatorUsers struct {
	global.GVA_MODEL
	AuthorityId *int   `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:用户角色ID;"`
	Email       string `json:"email" form:"email" gorm:"column:email;comment:用户邮箱;size:191;"`
	Enable      *int   `json:"enable" form:"enable" gorm:"column:enable;comment:用户是否被冻结 1正常 2冻结;"`
	HeaderImg   string `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:用户头像;size:191;"`
	NickName    string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;size:191;"`
	Password    string `json:"password" form:"password" gorm:"column:password;comment:用户登录密码;size:191;"`
	Phone       string `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号;size:191;"`
	Username    string `json:"username" form:"username" gorm:"column:username;comment:用户登录名;size:191;"`
	Uuid        string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户UUID;size:191;"`
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName CsOperatorUsers 表名
func (CsOperatorUsers) TableName() string {
	return "cs_operator_users"
}
