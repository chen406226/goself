package mysqlDb

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID    `json:"uuid"`
	Username    string       `json:"userName"`
	Password    string       `json:"-"`
	NickName    string       `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg   string       `json:"headerImg" gorm:"default:'http://qmplusimg.henrongyi.top/1571627762timg.jpg'"`
	Authority   SysAuthority `json:"authority" gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId"`
	AuthorityId string       `json:"authorityId" gorm:"default:888"`
}
