package operator

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operator"
	operatorReq "github.com/flipped-aurora/gin-vue-admin/server/model/operator/request"
)

type CsOperatorAuthoritiesService struct {
}

// CreateCsOperatorAuthorities 创建CsOperatorAuthorities记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorAuthoritiesService *CsOperatorAuthoritiesService) CreateCsOperatorAuthorities(csOperatorAuthorities *operator.CsOperatorAuthorities) (err error) {
	err = global.GVA_DB.Create(csOperatorAuthorities).Error
	return err
}

// DeleteCsOperatorAuthorities 删除CsOperatorAuthorities记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorAuthoritiesService *CsOperatorAuthoritiesService) DeleteCsOperatorAuthorities(csOperatorAuthorities operator.CsOperatorAuthorities) (err error) {
	err = global.GVA_DB.Delete(&csOperatorAuthorities).Error
	return err
}

// DeleteCsOperatorAuthoritiesByIds 批量删除CsOperatorAuthorities记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorAuthoritiesService *CsOperatorAuthoritiesService) DeleteCsOperatorAuthoritiesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]operator.CsOperatorAuthorities{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCsOperatorAuthorities 更新CsOperatorAuthorities记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorAuthoritiesService *CsOperatorAuthoritiesService) UpdateCsOperatorAuthorities(csOperatorAuthorities operator.CsOperatorAuthorities) (err error) {
	err = global.GVA_DB.Save(&csOperatorAuthorities).Error
	return err
}

// GetCsOperatorAuthorities 根据id获取CsOperatorAuthorities记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorAuthoritiesService *CsOperatorAuthoritiesService) GetCsOperatorAuthorities(id uint) (csOperatorAuthorities operator.CsOperatorAuthorities, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&csOperatorAuthorities).Error
	return
}

// GetCsOperatorAuthoritiesInfoList 分页获取CsOperatorAuthorities记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorAuthoritiesService *CsOperatorAuthoritiesService) GetCsOperatorAuthoritiesInfoList(info operatorReq.CsOperatorAuthoritiesSearch) (list []operator.CsOperatorAuthorities, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&operator.CsOperatorAuthorities{})
	var csOperatorAuthoritiess []operator.CsOperatorAuthorities
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

	err = db.Find(&csOperatorAuthoritiess).Error
	return csOperatorAuthoritiess, total, err
}
