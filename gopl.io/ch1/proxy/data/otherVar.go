package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"io/ioutil"
	"os/exec"
	"strings"
	"syscall"
)


type Config struct {
	ProviderList   	[]User  	`json:"providerList"`
	PublishList   	[]Pub  		`json:"publishList"`
	Default			string		`json:"default"`
	PubName			string		`json:"pubName"`
	SourceFolder	string		`json:"sourceFolder"`
}

type User struct {
	Name 	string	`json:"name"`
	Host	string	`json:"host"`
}
type Pub struct {
	Name 	string	`json:"name"`
	IP	string	`json:"ip"`
	User	string	`json:"user"`
	Pass	string	`json:"pass"`
	ToDir	string  `json:"toDir"`
	Build	string  `json:"build"`
}

var(
	CashData = &Config{
		Default: "Dev",
		ProviderList: []User {
			User{Name: "Dev",Host: "http://10.10.0.123:8100"},
			User{Name: "Local",Host: "http://localhost:8088"},
			User{Name: "Test",Host: "http://10.10.0.123:8088"},
			User{Name: "ZhangZeng",Host: "http://10.10.0.43:8088"},
			User{Name: "HouZhiYong",Host: "http://10.10.0.76:8200"},
		},
		PublishList: []Pub {
			{Name: "PC测试",
				IP: "10.10.0.123",
				User: "Administrator",
				Pass: "Nc@test",
				ToDir: "D\\public\\NcManageUI",
				Build: "npm run build:test"},
		},
		PubName: "PC测试",
		SourceFolder : "",
	}
	RunName = "Dev"
	PubName = "PC测试"
	SourceFolder = ""
)

func init() {
	con := LoadConfig()
	if con != nil {
		CashData = con
		RunName = CashData.Default
		SourceFolder = CashData.SourceFolder
	} else {
		SaveConfig(CashData)
	}
}

var configPath = "C:/Users/Administrator/proxyConfig.json"

func SaveConfig(config *Config){

	data,err:=json.Marshal(config)
	if err!=nil{
		//log.Fatal(err)
	}
	err=ioutil.WriteFile(configPath,data,0660)
	if err!=nil{
		//log.Fatal(err)
	}

}

func SaveDefault()  {
	SaveConfig(CashData)
}

func GetDefault()  {

}

func GetPub() (s Pub, err error) {
	for _, v := range CashData.PublishList {
		if v.Name == CashData.PubName {
			return v,nil
		}
	}
	err = errors.New("not find")
	return s, err
}

func FirstConnection(win fyne.Window, lb *canvas.Text)  {
	pub, errs:= GetPub()
	if errs != nil {
		dialog.ShowError(errors.New("未查找到发布服务信息"), win)
		return
	}
	if lb != nil {
		lb.Text = "连接中..."
		lb.Refresh()
	}
	cmd := exec.Command("cmd.exe", "/c", "ping "+pub.IP)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err != nil {
		dialog.ShowError(errors.New("发布服务ip查找失败"), win)
		return
	}
	if lb != nil {
		lb.Text = "If The Window Opens Successfully , Close To Continue !!!"
		lb.Refresh()
	}
	//
	//cmd = exec.Command("cmd.exe", "/c", "mstsc /v: 10.10.0.123 /console")
	//err = cmd.Run()
	//cmd = exec.Command("cmd.exe", "/c", "net use \\\\10.10.0.123\\ipc$ Nc@test /user:Administrator")
	cmd = exec.Command("cmd.exe", "/c", "net use \\\\"+pub.IP+"\\ipc$ "+pub.Pass+" /user:"+pub.User)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = cmd.Run()
	if win == nil {
		return
	}
	if err != nil {
		fmt.Println("Net Use Error")
		dialog.ShowError(errors.New("发布服务用户名或密码错误"), win)
		return
	}
	dialog.ShowInformation("Success", "发布服务连接成功", win)
}

func MoveFile(win fyne.Window, lb *canvas.Text)  {
	pub, errs:= GetPub()
	if errs != nil {
		dialog.ShowError(errors.New("未查找到发布服务信息"), win)
		return
	}
	t := strings.Replace(CashData.SourceFolder,"/","\\",100)
	buildFolder := CashData.SourceFolder

	lb.Text = "pull code from origin/master"
	lb.Refresh()

	cmd := exec.Command("cmd.exe", "/c", "cd/d "+ buildFolder +" && git pull")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err != nil {
		dialog.ShowError(errors.New("pull code Error"), win)
		return
	}

	lb.Text = "Build Project ..."
	lb.Refresh()

	cmd = exec.Command("cmd.exe", "/c", "cd/d "+ buildFolder +" && "+pub.Build)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = cmd.Run()
	if err != nil {
		dialog.ShowError(errors.New("Build Error"), win)
		return
	}
	lb.Text = "Send File ..."
	lb.Refresh()
	cmd = exec.Command("cmd.exe", "/c", "xcopy "+t+"\\dist \\\\"+pub.IP+"\\"+pub.ToDir+" /s/e/y")
	//cmd = exec.Command("cmd.exe", "/c", "xcopy "+t+"\\src\\Presentation\\NCManageUI\\dist \\\\10.10.0.123\\D\\public\\NcManageUI /s/e/y")
	//cmd = exec.Command("cmd.exe", "/c", "xcopy "+t+"\\src\\Presentation\\NCManageUI\\dist \\\\10.10.0.123\\D\\test /s/e/y")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = cmd.Run()
	if err != nil {
		dialog.ShowError(errors.New("Publish File Error,Please Try To Click 'Connect Teleservice' Or Check Your File Sharing Has Been Set Up"), win)
		return
	}
	dialog.ShowInformation("Success", "Publish new Version Success", win)
}


func LoadConfig()(config *Config){
	data,err:=ioutil.ReadFile(configPath)
	if err!=nil{
		//log.Println(err)
		return nil
	}
	//for _, line := range strings.Split(string(data), "\n") {
	//    fmt.Println(line)
	//}
	config =&Config{}
	err=json.Unmarshal(data,&config)
	if err!=nil{
		//log.Println(err)
	}
	return config
}