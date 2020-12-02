package request

import "student/model/mysqlDb"

type SysOperationRecordSearch struct {
	mysqlDb.SysOperationRecord
	PageInfo
}