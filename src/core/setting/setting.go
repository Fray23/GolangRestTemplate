package setting

import (
	"log"
	"github.com/go-ini/ini"
)


type App struct {
	Port      string
	JwtSecret string
}

type Database struct {
	Host        string
	User        string
	Password    string
	Name        string
	Port        string
}

var AppSetting = &App{}
var DatabaseSetting = &Database{}

var cfg *ini.File

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

func Setup() {
	var err error
	cfg, err = ini.Load("settings.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("db", DatabaseSetting)
}
