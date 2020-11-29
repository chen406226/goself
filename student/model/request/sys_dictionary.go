package request

import "student/model/mysqlDb"

type SysDictionarySearch struct {
	mysqlDb.SysDictionary
	PageInfo
}