package tenant

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tenant"
	tenantReq "github.com/flipped-aurora/gin-vue-admin/server/model/tenant/request"
	"gorm.io/gorm"
)

type CsTenantUserService struct {
}

// CreateCsTenantUser 创建CsTenantUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantUserService *CsTenantUserService) CreateCsTenantUser(csTenantUser *tenant.CsTenantUser) (err error) {
	err = global.GVA_DB.Create(csTenantUser).Error
	return err
}

// DeleteCsTenantUser 删除CsTenantUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantUserService *CsTenantUserService) DeleteCsTenantUser(csTenantUser tenant.CsTenantUser) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&tenant.CsTenantUser{}).Where("id = ?", csTenantUser.ID).Update("deleted_by", csTenantUser.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&csTenantUser).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCsTenantUserByIds 批量删除CsTenantUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantUserService *CsTenantUserService) DeleteCsTenantUserByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&tenant.CsTenantUser{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&tenant.CsTenantUser{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCsTenantUser 更新CsTenantUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantUserService *CsTenantUserService) UpdateCsTenantUser(csTenantUser tenant.CsTenantUser) (err error) {
	err = global.GVA_DB.Save(&csTenantUser).Error
	return err
}

// GetCsTenantUser 根据id获取CsTenantUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantUserService *CsTenantUserService) GetCsTenantUser(id uint) (csTenantUser tenant.CsTenantUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&csTenantUser).Error
	return
}

// GetCsTenantUserInfoList 分页获取CsTenantUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantUserService *CsTenantUserService) GetCsTenantUserInfoList(info tenantReq.CsTenantUserSearch) (list []tenant.CsTenantUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&tenant.CsTenantUser{})
	var csTenantUsers []tenant.CsTenantUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+info.NickName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&csTenantUsers).Error
	return csTenantUsers, total, err
}
