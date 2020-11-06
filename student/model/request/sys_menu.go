package request

import "student/model/mysqlDb"

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus		[]mysqlDb.SysBaseMenu
	AuthorityId	string
}

// Get role by id structure
type AuthorityIdInfo struct {
	AuthorityId string
}
