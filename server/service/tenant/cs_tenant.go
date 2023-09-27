package tenant

import (
	"errors"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tenant"
	tenantReq "github.com/flipped-aurora/gin-vue-admin/server/model/tenant/request"
	system2 "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type CsTenantService struct {
}

func (csTenantService *CsTenantService) CreateTenantIDCasbin(tenantID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetCasbinName(tenantID)).AutoMigrate(&adapter.CasbinRule{})
	return err
}

func (csTenantService *CsTenantService) CreateTenantIDAuthMenu(tenantID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetAuthMenuTableName(tenantID)).AutoMigrate(&system.SysAuthorityMenu{})
	return err
}

func (csTenantService *CsTenantService) CreateTenetUserTableName(tenantID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetUserTableName(tenantID)).AutoMigrate(&system.SysUserTable{})
	return err
}

func (csTenantService *CsTenantService) CreateTenetAuthsTable(tenantID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetAuthsTable(tenantID)).AutoMigrate(&system.SysAuthorityTable{})
	return err
}

func (csTenantService *CsTenantService) CreateUserAuthorityTable(tenantID uint) (err error) {
	err = global.GVA_DB.Table(utils.GetUserAuthorityTableName(tenantID)).AutoMigrate(&system.SysUserAuthority{})
	return err
}

var sysUserServer = new(system2.UserService)
var sysAuthorityServer = new(system2.AuthorityService)

// CreateCsTenant 创建CsTenant记录
// Author [piexlmax](https://github.com/piexlmax)
func (csTenantService *CsTenantService) CreateCsTenant(csTenant *tenant.CsTenant) (err error) {

	if global.GVA_DB.First(&tenant.CsTenant{}, "username = ?", csTenant.Username).Error == nil {
		return errors.New("用户名已注册")
	}
	oldPassWord := csTenant.Password
	csTenant.OnlyKey, _ = uuid.NewV4()
	csTenant.Password = utils.BcryptHash(csTenant.Password)
	err = global.GVA_DB.Create(csTenant).Error
	if err != nil {
		return err
	}

	err = csTenantService.CreateTenetUserTableName(csTenant.ID)
	if err != nil {
		return err
	}

	err = csTenantService.CreateTenetAuthsTable(csTenant.ID)
	if err != nil {
		return err
	}

	err = csTenantService.CreateUserAuthorityTable(csTenant.ID)
	if err != nil {
		return err
	}

	user := &system.SysUser{
		Username: csTenant.Username,
		Password: oldPassWord,
	}
	user.TenantID = csTenant.ID
	user.NickName = "超级管理员"
	user.HeaderImg = csTenant.Logo
	user.AuthorityId = 888
	_, err = sysUserServer.Register(*user, csTenant.ID)
	if err != nil {
		return err
	}

	auth := &system.SysAuthority{
		AuthorityId:   888,
		AuthorityName: "超级管理员",
		ParentId:      utils.Pointer(uint(0)),
		DefaultRouter: "dashboard",
	}

	_, err = sysAuthorityServer.CreateAuthority(*auth, csTenant.ID)
	if err != nil {
		return err
	}

	err = csTenantService.CreateTenantIDCasbin(csTenant.ID)
	if err != nil {
		return err
	}
	err = csTenantService.CreateTenantIDAuthMenu(csTenant.ID)
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

func (csTenantService *CsTenantService) GetTenantApis(id string) (csTenantApis []tenant.CsTenantApis, err error) {
	err = global.GVA_DB.Where("tenant_id = ?", id).Find(&csTenantApis).Error
	return
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

func (csTenantService *CsTenantService) GetTenantMenus(id string) (csTenantMenus []tenant.CsTenantMenus, err error) {
	err = global.GVA_DB.Where("tenant_id = ?", id).Find(&csTenantMenus).Error
	return
}

func (csTenantService *CsTenantService) GetTenantIDByOnlyKey(onlyKey string) uint {
	var csTenant tenant.CsTenant
	global.GVA_DB.Where("only_key = ?", onlyKey).First(&csTenant)
	return csTenant.ID
}
