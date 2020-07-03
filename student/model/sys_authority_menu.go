package model

import "student/model/mysqlDb"

type SysMenu struct {
	mysqlDb.SysBaseMenu
	MenuId      string    `json:"menuId"`
	AuthorityId string    `json:"-"`
	Children    []SysMenu `json:"children"`
}