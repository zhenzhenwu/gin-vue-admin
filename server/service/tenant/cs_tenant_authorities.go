package tenant

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tenant"
	tenantReq "github.com/flipped-aurora/gin-vue-admin/server/model/tenant/request"
	"gorm.io/gorm"
)

type CsTenantAuthoritiesService struct {
}

// CreateCsTenantAuthorities 创建CsTenantAuthorities记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantAuthoritiesService *CsTenantAuthoritiesService) CreateCsTenantAuthorities(csTenantAuthorities *tenant.CsTenantAuthorities) (err error) {
	err = global.GVA_DB.Create(csTenantAuthorities).Error
	return err
}

// DeleteCsTenantAuthorities 删除CsTenantAuthorities记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantAuthoritiesService *CsTenantAuthoritiesService) DeleteCsTenantAuthorities(csTenantAuthorities tenant.CsTenantAuthorities) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&tenant.CsTenantAuthorities{}).Where("id = ?", csTenantAuthorities.ID).Update("deleted_by", csTenantAuthorities.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&csTenantAuthorities).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCsTenantAuthoritiesByIds 批量删除CsTenantAuthorities记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantAuthoritiesService *CsTenantAuthoritiesService) DeleteCsTenantAuthoritiesByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&tenant.CsTenantAuthorities{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&tenant.CsTenantAuthorities{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCsTenantAuthorities 更新CsTenantAuthorities记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantAuthoritiesService *CsTenantAuthoritiesService) UpdateCsTenantAuthorities(csTenantAuthorities tenant.CsTenantAuthorities) (err error) {
	err = global.GVA_DB.Save(&csTenantAuthorities).Error
	return err
}

// GetCsTenantAuthorities 根据id获取CsTenantAuthorities记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantAuthoritiesService *CsTenantAuthoritiesService) GetCsTenantAuthorities(id uint) (csTenantAuthorities tenant.CsTenantAuthorities, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&csTenantAuthorities).Error
	return
}

// GetCsTenantAuthoritiesInfoList 分页获取CsTenantAuthorities记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantAuthoritiesService *CsTenantAuthoritiesService) GetCsTenantAuthoritiesInfoList(info tenantReq.CsTenantAuthoritiesSearch) (list []tenant.CsTenantAuthorities, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&tenant.CsTenantAuthorities{})
	var csTenantAuthoritiess []tenant.CsTenantAuthorities
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&csTenantAuthoritiess).Error
	return csTenantAuthoritiess, total, err
}
