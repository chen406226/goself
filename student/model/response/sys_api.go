package response

import "student/model/mysqlDb"

type SysAPIResponse struct {
	Api	mysqlDb.SysApi	`json:"api"`
}

type SysAPIListResponse struct {
	Apis	[]mysqlDb.SysApi	`json:"apis"`
}