package response

import "student/model/mysqlDb"

type SysAuthorityResponse struct {
	Authority	mysqlDb.SysAuthority	`json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority				mysqlDb.SysAuthority	`json:"authority"`
	OldAuthorityId	string								`json:"oldAuthorityId"`
}