package main

import (
	"net/http"
	_ "net/http/pprof"

	"go-api-jsonrssfeed/api"
	util "git.aventer.biz/AVENTER/util"
	"github.com/sirupsen/logrus"
)

// MinVersion is the build version
var MinVersion string

func main() {
	util.SetLogging(config.LogLevel, config.EnableSyslog, config.AppName)
	logrus.Println(config.AppName+" build"+MinVersion, config.APIBind, config.APIPort)

	api.SetConfig(config)

	http.Handle("/", api.Commands())

	if err := http.ListenAndServe(config.APIBind+":"+config.APIPort, nil); err != nil {
		logrus.Fatalln("ListenAndServe: ", err)
	}	
}
