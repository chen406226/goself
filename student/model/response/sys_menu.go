package response

import (
	"student/model/mysqlDb"
)

type SysMenusResponse struct {
	Menus []mysqlDb.SysMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []mysqlDb.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu mysqlDb.SysBaseMenu `json:"menu"`
}
