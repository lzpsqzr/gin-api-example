package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	RunEnv string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	Loadenvconf()
}

func Loadenvconf() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	RunEnv = Cfg.Section("").Key("RUN_ENV").MustString("local")
	fmt.Println(RunEnv)
	if RunEnv == "local" {
		sec, err := Cfg.GetSection("local")
		if err != nil {
			log.Fatalf("Fail to get section 'server': %v", err)
		}

		HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
		ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
		WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
		JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
		PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	} else if RunEnv == "dev" {
		sec, err := Cfg.GetSection("local")
		if err != nil {
			log.Fatalf("Fail to get section 'server': %v", err)
		}

		HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
		ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
		WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
		JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
		PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	} else {
		log.Fatalf("Wrong run mode , please check the runmode in config file !")
	}

}
