package data

import (
	"encoding/json"
	"io/ioutil"
)


type Config struct {
	ProviderList   	[]User  	`json:"providerList"`
	Default					string							`json:"default"`
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
	}
	RunName = "Dev"
)

func init() {
	con := LoadConfig()
	if con != nil {
		CashData = con
		RunName = CashData.Default
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