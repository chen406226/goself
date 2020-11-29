package request

import "student/model/mysqlDb"

type SysDictionaryDetailSearch struct {
	mysqlDb.SysDictionaryDetail
	PageInfo
}