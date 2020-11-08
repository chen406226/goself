package service

import (
	"errors"
	"fmt"
	"strconv"
	"student/global"
	"student/model/mysqlDb"
)

// @title    AddBaseMenu
// @description   函数的详细描述
// @auth                     （2020/04/05  20:22）
// @param     menu            *model.SysBaseMenu
// @return    err             error
// 增加基础路由
func AddBaseMenu(u mysqlDb.SysBaseMenu) (err error) {
	var menu mysqlDb.SysBaseMenu
	//判断是否存在
	//if !errors.Is(global.GL_DB.Where("name = ?", menu.Name).First(&mysqlDb.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
	//	err = errors.New("存在重复name，请修改name")
	//}
	notRegister := global.GL_DB.Where("name = ?",menu.Name).First(&menu).RecordNotFound()
	if !notRegister {
		return errors.New("存在重复name,请修改name")
	} else {
		err = global.GL_DB.Create(&u).Error
	}
	return err
}

// @title   getMenuTreeMap
// @description    获取路由总树map
// @auth       qm      (2020/05/06 10:26)
// @return     err             error
// @return    menusMsp            map{string}[]SysBaseMenu
func getMenuTreeMap(authorityId string) (err error,treeMap map[string][]mysqlDb.SysMenu) {
	var allMenus []mysqlDb.SysMenu
	treeMap = make(map[string][]mysqlDb.SysMenu)
	err = global.GL_DB.Where("authority_id = ?",authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		fmt.Println(v.ParentId)
		treeMap[v.ParentId] = append(treeMap[v.ParentId],v)
	}
	return err,treeMap
}

// @title    GetMenuTree
// @description   获取动态菜单树
// @auth                     （2020/04/05  20:22）
// @param     authorityId     string
// @return    err             error
// @return    menus           []model.SysMenu
func GetMenuTree(authorityId string) (err error,menus []mysqlDb.SysMenu) {
	err, menuTree := getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}


// @title    getChildrenList
// @description   获取子菜单
// @auth                     （2020/04/05  20:22）
// @param     menu            *model.SysMenu
// @param     sql             string
// @return    err             error
func getChildrenList(menu *mysqlDb.SysMenu,treeMap map[string][]mysqlDb.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i],treeMap)
	}
	return err
}


// @title    GetMenuList
// @description   获取路由分页
// @auth                     （2020/04/05  20:22）
// @param     info            request.PageInfo
// @return    err             error
// @return    list            interface{}
// @return    total           int
func GetMenuList() (err error,list interface{},total int64) {
	var menuList []mysqlDb.SysBaseMenu
	err, treeMap := getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i< len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i],treeMap)
	}
	return err,menuList, total
}

// @title    getBaseChildrenList
// @description   get children of menu, 获取菜单的子菜单
// @auth                     （2020/04/05  20:22）
// @param     menu            *model.SysBaseMenu
// @return    err             error

func getBaseChildrenList(menu *mysqlDb.SysBaseMenu, treeMap map[string][]mysqlDb.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// @title   getBaseMenuTreeMap
// @description    获取路由总树map
// @auth       qm      (2020/05/06 10:26)
// @return     err             error
// @return    menusMsp            map{string}[]SysBaseMenu

func getBaseMenuTreeMap() (err error, treeMap map[string][]mysqlDb.SysBaseMenu) {
	var allMenus []mysqlDb.SysBaseMenu
	treeMap = make(map[string][]mysqlDb.SysBaseMenu)
	err = global.GL_DB.Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}
