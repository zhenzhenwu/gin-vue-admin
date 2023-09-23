package utils

import "fmt"

func GetCasbinName(tenantID uint) string {
	if tenantID != 0 {
		return fmt.Sprintf("casbin_rule_tenant_%d", tenantID)
	}
	return "casbin_rule"
}

func GetAuthMenuTableName(tenantID uint) string {
	if tenantID != 0 {
		return fmt.Sprintf("cs_authority_menus_tenant_%d", tenantID)
	}
	return "sys_authority_menus"
}

func GetUserTableName(tenantID uint) string {
	if tenantID != 0 {
		return fmt.Sprintf("cs_user_tenant_%d", tenantID)
	}
	return "sys_users"
}

func GetAuthsTable(tenantID uint) string {
	if tenantID != 0 {
		return fmt.Sprintf("cs_authorities_tenant_%d", tenantID)
	}
	return "sys_authorities"
}
