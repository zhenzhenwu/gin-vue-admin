package tenant

type CsTenantApis struct {
	ID       uint `gorm:"primarykey"` // 主键ID
	TenantId uint `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户id;"`
	ApiId    uint `json:"apiId" form:"apiId" gorm:"column:api_id;comment:api id;"`
}

// TableName CsTenant 表名
func (CsTenantApis) TableName() string {
	return "cs_tenant_apis"
}
