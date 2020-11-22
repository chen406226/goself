package service

import (
	"errors"
	"github.com/casbin/casbin/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	//"github.com/casbin/casbin"
	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	//gormadapter "github.com/casbin/gorm-adapter"
	"strings"
	"student/global"
	"student/model/mysqlDb"
	"student/model/request"
)

// @title    Casbin
// @description   store to DB, 持久化到数据库  引入自定义规则
// @auth                     （2020/04/05  20:22）

func Casbin() *casbin.Enforcer {
	admin := global.GL_CONFIG.Mysql
	a, _ := gormadapter.NewAdapter(global.GL_CONFIG.System.DbType, admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname, true)
	//a := gormadapter.NewAdapterByDB(global.GL_DB)
	e, _ := casbin.NewEnforcer(global.GL_CONFIG.Casbin.ModelPath,a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = e.LoadPolicy()
	return e
}


// @title    ParamsMatch
// @description   customized rule, 自定义规则函数
// @auth                     （2020/04/05  20:22）
// @param     fullNameKey1    string
// @param     key2            string
// @return                    bool

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	//剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// @title    ParamsMatchFunc
// @description   customized function, 自定义规则函数
// @auth                     （2020/04/05  20:22）
// @param     args            ...interface{}
// @return                    interface{}
// @return                    error

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(ParamsMatch(name1, name2)), nil
}

// @title    GetPolicyPathByAuthorityId
// @description   get policy path by authorityId, 获取权限列表
// @auth                     （2020/04/05  20:22）
// @param     authorityId     string
// @return                    []string
func GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := Casbin()
	list := e.GetFilteredPolicy(0,authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps,request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return  pathMaps
}

// @title    UpdateCasbin
// @description   update casbin authority, 更新casbin权限
// @auth                     （2020/04/05  20:22）
// @param     authorityId      string
// @param     casbinInfos      []CasbinInfo
// @return                     error

func UpdateCasbin(authorityId string,casbinInfos []request.CasbinInfo) error {
	ClearCasbin(0,authorityId)
	var rules [][]string
	for _, v := range casbinInfos {
		cm := mysqlDb.CasbinModel{
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules,[]string{cm.AuthorityId, cm.Path,cm.Method})
	}
	e := Casbin()
	success,_ := e.AddPolicies(rules)
	if success == false {
		return errors.New("存在相同api,添加失败，请联系管理员")
	}
	return nil
}

// @title    ClearCasbin
// @description   清除匹配的权限
// @auth                     （2020/04/05  20:22）
// @param     v               int
// @param     p               string
// @return                    bool

func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	success,_ := e.RemoveFilteredPolicy(v, p...)
	return success

}

// @title    UpdateCasbinApi
// @description   update casbin apis, API更新随动
// @auth                     （2020/04/05  20:22）
// @param     oldPath          string
// @param     newPath          string
// @param     oldMethod        string
// @param     newMethod        string
// @return                     error

func UpdateCasbinApi(oldPath string,newPath string,oldMethod string,newMethod string) error {
	err := global.GL_DB.Table("casbin_rule").Model(&mysqlDb.CasbinModel{}).
		Where("v1 = ? AND v2 =?",oldPath,oldMethod).
		Updates(map[string]interface{}{"v1": newPath,"v2": newMethod}).Error
	return err
}






