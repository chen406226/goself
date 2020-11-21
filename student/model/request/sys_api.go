package request

import "student/model/mysqlDb"

type SearchApiParams struct {
	mysqlDb.SysApi
	PageInfo
	OrderKey	string	`json:"orderKey"`
	Desc			bool		`json:"desc"`
}