package operator

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operator"
	operatorReq "github.com/flipped-aurora/gin-vue-admin/server/model/operator/request"
	"gorm.io/gorm"
)

type CsOperatorUsersService struct {
}

// CreateCsOperatorUsers 创建CsOperatorUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorUsersService *CsOperatorUsersService) CreateCsOperatorUsers(csOperatorUsers *operator.CsOperatorUsers) (err error) {
	err = global.GVA_DB.Create(csOperatorUsers).Error
	return err
}

// DeleteCsOperatorUsers 删除CsOperatorUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorUsersService *CsOperatorUsersService) DeleteCsOperatorUsers(csOperatorUsers operator.CsOperatorUsers) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&operator.CsOperatorUsers{}).Where("id = ?", csOperatorUsers.ID).Update("deleted_by", csOperatorUsers.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&csOperatorUsers).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCsOperatorUsersByIds 批量删除CsOperatorUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorUsersService *CsOperatorUsersService) DeleteCsOperatorUsersByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&operator.CsOperatorUsers{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&operator.CsOperatorUsers{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCsOperatorUsers 更新CsOperatorUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorUsersService *CsOperatorUsersService) UpdateCsOperatorUsers(csOperatorUsers operator.CsOperatorUsers) (err error) {
	err = global.GVA_DB.Save(&csOperatorUsers).Error
	return err
}

// GetCsOperatorUsers 根据id获取CsOperatorUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorUsersService *CsOperatorUsersService) GetCsOperatorUsers(id uint) (csOperatorUsers operator.CsOperatorUsers, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&csOperatorUsers).Error
	return
}

// GetCsOperatorUsersInfoList 分页获取CsOperatorUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorUsersService *CsOperatorUsersService) GetCsOperatorUsersInfoList(info operatorReq.CsOperatorUsersSearch) (list []operator.CsOperatorUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&operator.CsOperatorUsers{})
	var csOperatorUserss []operator.CsOperatorUsers
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

	err = db.Find(&csOperatorUserss).Error
	return csOperatorUserss, total, err
}
