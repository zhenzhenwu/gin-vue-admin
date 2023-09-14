package operator

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operator"
	operatorReq "github.com/flipped-aurora/gin-vue-admin/server/model/operator/request"
	"gorm.io/gorm"
)

type CsOperatorService struct {
}

// CreateCsOperator 创建CsOperator记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorService *CsOperatorService) CreateCsOperator(csOperator *operator.CsOperator) (err error) {
	err = global.GVA_DB.Create(csOperator).Error
	return err
}

// DeleteCsOperator 删除CsOperator记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorService *CsOperatorService) DeleteCsOperator(csOperator operator.CsOperator) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&operator.CsOperator{}).Where("id = ?", csOperator.ID).Update("deleted_by", csOperator.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&csOperator).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCsOperatorByIds 批量删除CsOperator记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorService *CsOperatorService) DeleteCsOperatorByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&operator.CsOperator{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&operator.CsOperator{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCsOperator 更新CsOperator记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorService *CsOperatorService) UpdateCsOperator(csOperator operator.CsOperator) (err error) {
	err = global.GVA_DB.Save(&csOperator).Error
	return err
}

// GetCsOperator 根据id获取CsOperator记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorService *CsOperatorService) GetCsOperator(id uint) (csOperator operator.CsOperator, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&csOperator).Error
	return
}

// GetCsOperatorInfoList 分页获取CsOperator记录
// Author [piexlmax](https://github.com/piexlmax)
func (csOperatorService *CsOperatorService) GetCsOperatorInfoList(info operatorReq.CsOperatorSearch) (list []operator.CsOperator, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&operator.CsOperator{})
	var csOperators []operator.CsOperator
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

	err = db.Find(&csOperators).Error
	return csOperators, total, err
}

func (csOperatorService *CsOperatorService) SetOperatorApis(operatorApis operatorReq.CsOperatorApisReq) error {
	var csOperatorApis []operator.CsOperatorApis
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err := global.GVA_DB.Where("operator_id = ?", operatorApis.OperatorId).Delete(&csOperatorApis).Error
		if err != nil {
			return err
		}
		for _, apiId := range operatorApis.ApiIDs {
			csOperatorApis = append(csOperatorApis, operator.CsOperatorApis{
				OperatorId: operatorApis.OperatorId,
				ApiId:      apiId,
			})
		}
		return global.GVA_DB.Create(&csOperatorApis).Error
	})
}

func (csOperatorService *CsOperatorService) SetOperatorMenus(operatorMenus operatorReq.CsOperatorMenusReq) error {
	var csOperatorMenus []operator.CsOperatorMenus
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err := global.GVA_DB.Where("operator_id = ?", operatorMenus.OperatorId).Delete(&csOperatorMenus).Error
		if err != nil {
			return err
		}
		for _, menuID := range operatorMenus.MenuIDs {
			csOperatorMenus = append(csOperatorMenus, operator.CsOperatorMenus{
				OperatorId: operatorMenus.OperatorId,
				MenuId:     menuID,
			})
		}
		return global.GVA_DB.Create(&csOperatorMenus).Error
	})
}
