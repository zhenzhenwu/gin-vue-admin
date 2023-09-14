package tenant

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tenant"
	tenantReq "github.com/flipped-aurora/gin-vue-admin/server/model/tenant/request"
	"gorm.io/gorm"
)

type CsTenantService struct {
}

// CreateCsTenant 创建CsTenant记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantService *CsTenantService) CreateCsTenant(csTenant *tenant.CsTenant) (err error) {
	err = global.GVA_DB.Create(csTenant).Error
	return err
}

// DeleteCsTenant 删除CsTenant记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantService *CsTenantService) DeleteCsTenant(csTenant tenant.CsTenant) (err error) {
	err = global.GVA_DB.Delete(&csTenant).Error
	return err
}

// DeleteCsTenantByIds 批量删除CsTenant记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantService *CsTenantService) DeleteCsTenantByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]tenant.CsTenant{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCsTenant 更新CsTenant记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantService *CsTenantService) UpdateCsTenant(csTenant tenant.CsTenant) (err error) {
	err = global.GVA_DB.Save(&csTenant).Error
	return err
}

// GetCsTenant 根据id获取CsTenant记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantService *CsTenantService) GetCsTenant(id uint) (csTenant tenant.CsTenant, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&csTenant).Error
	return
}

// GetCsTenantInfoList 分页获取CsTenant记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantService *CsTenantService) GetCsTenantInfoList(info tenantReq.CsTenantSearch) (list []tenant.CsTenant, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&tenant.CsTenant{})
	var csTenants []tenant.CsTenant
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	if info.StartEndTime != nil && info.EndEndTime != nil {
		db = db.Where("end_time BETWEEN ? AND ? ", info.StartEndTime, info.EndEndTime)
	}
	if info.Allocation != nil {
		db = db.Where("allocation = ?", info.Allocation)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&csTenants).Error
	return csTenants, total, err
}

func (csTenantService *CsTenantService) SetTenantApis(tenantApis tenantReq.CsTenantApisReq) error {
	var csTenantApis []tenant.CsTenantApis
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err := global.GVA_DB.Where("tenant_id = ?", tenantApis.TenantId).Delete(&csTenantApis).Error
		if err != nil {
			return err
		}
		for _, apiId := range tenantApis.ApiIDs {
			csTenantApis = append(csTenantApis, tenant.CsTenantApis{
				TenantId: tenantApis.TenantId,
				ApiId:    apiId,
			})
		}
		return global.GVA_DB.Create(&csTenantApis).Error
	})
}

func (csTenantService *CsTenantService) SetTenantMenus(tenantMenus tenantReq.CsTenantMenusReq) error {
	var csTenantMenus []tenant.CsTenantMenus
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err := global.GVA_DB.Where("tenant_id = ?", tenantMenus.TenantId).Delete(&csTenantMenus).Error
		if err != nil {
			return err
		}
		for _, menuID := range tenantMenus.MenuIDs {
			csTenantMenus = append(csTenantMenus, tenant.CsTenantMenus{
				TenantId: tenantMenus.TenantId,
				MenuId:   menuID,
			})
		}
		return global.GVA_DB.Create(&csTenantMenus).Error
	})
}
