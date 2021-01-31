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
	"time"
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
	err := cmd.Run()
	if err != nil {
		dialog.ShowError(errors.New("IP Cannot Connect"), win)
		return
	}

	lb.Text = "Login Teleservice ..."
	lb.Refresh()
	cmd = exec.Command("cmd.exe", "/c", "mstsc /v: 10.10.0.123 /console")
	err = cmd.Run()
	if err != nil {
		dialog.ShowError(errors.New("Connection Error"), win)
		return
	}
	cmd = exec.Command("cmd.exe", "/c", "net use \\\\10.10.0.123\\ipc$ Nc@test /user:Administrator")
	err = cmd.Run()
	if err != nil {
		dialog.ShowError(errors.New("Uesr Or PassWord Error"), win)
	}
}

func MoveFile(win fyne.Window, lb *canvas.Text)  {
	t := strings.Replace(CashData.SourceFolder,"/","\\",100)
	lb.Text = "Build Project ..."
	lb.Refresh()

	//cmd := exec.Command("cmd.exe", "/c", "cd/d "+ CashData.SourceFolder +" && npm run build")
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("NNNNNN")
	//
	//	dialog.ShowError(errors.New("Build Error"), win)
	//	return
	//}
	time.Sleep(time.Second *4)
	fmt.Println("WWWWWWWWWWWW")
	lb.Text = "Send File ..."
	lb.Refresh()
	time.Sleep(time.Second *4)

	cmd := exec.Command("cmd.exe", "/c", "xcopy "+t+"\\dist \\\\10.10.0.123\\D\\public\\NcManageUI /s/e/y")
	err := cmd.Run()
	if err != nil {
		fmt.Println("NNNNNNNNNNN")
		dialog.ShowError(errors.New("Publish File Error,Please Try To Click 'Connect Teleservice' Or Check Your File Sharing Has Been Set Up"), win)
		return
	}
	fmt.Println("WWWWWWW")
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