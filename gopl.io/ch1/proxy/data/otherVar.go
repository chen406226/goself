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
	Default					string		`json:"default"`
	SourceFolder		string		`json:"sourceFolder"`
}

type User struct {
	Name 	string	`json:"name"`
	Host	string	`json:"host"`
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
		SourceFolder : "",
	}
	RunName = "Dev"
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

func FirstConnection(win fyne.Window, lb *canvas.Text)  {
	lb.Text = "Connect Ip ..."
	lb.Refresh()
	cmd := exec.Command("cmd.exe", "/c", "ping 10.10.0.123")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err != nil {
		dialog.ShowError(errors.New("IP Cannot Connect"), win)
		return
	}

	lb.Text = "If The Window Opens Successfully , Close To Continue !!!"
	lb.Refresh()
	//
	//cmd = exec.Command("cmd.exe", "/c", "mstsc /v: 10.10.0.123 /console")
	//err = cmd.Run()
	//fmt.Println("%%%%%%%%%%%%%%%")
	//
	//if err != nil {
	//	dialog.ShowError(errors.New("Connection Error"), win)
	//	return
	//}
	cmd = exec.Command("cmd.exe", "/c", "net use \\\\10.10.0.123\\ipc$ Nc@test /user:Administrator")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = cmd.Run()
	if err != nil {
		fmt.Println("Net Use Error")
		dialog.ShowError(errors.New("Uesr Or PassWord Error"), win)
	}
	dialog.ShowInformation("Success", "Connection successful", win)
}

func MoveFile(win fyne.Window, lb *canvas.Text)  {
	t := strings.Replace(CashData.SourceFolder,"/","\\",100)
	buildFolder := CashData.SourceFolder + "/src/Presentation/NCManageUI"
	lb.Text = "Build Project ..."
	lb.Refresh()

	cmd := exec.Command("cmd.exe", "/c", "cd/d "+ buildFolder +" && npm run build")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err != nil {
		dialog.ShowError(errors.New("Build Error"), win)
		return
	}
	lb.Text = "Send File ..."
	lb.Refresh()
	cmd = exec.Command("cmd.exe", "/c", "xcopy "+t+"\\src\\Presentation\\NCManageUI\\dist \\\\10.10.0.123\\D\\public\\NcManageUI /s/e/y")
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