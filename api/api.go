package api

import (

	//	"encoding/json"

	"github.com/gorilla/mux"

	"net/http"

	cfg "go-api-jsonrssfeed/types"

	util "git.aventer.biz/AVENTER/util"
)

var config cfg.Config

// APIVersions give out a list of Versions
func APIVersions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "-")
	w.Write([]byte("/api/v0"))
}

// APIV0Version give out the version number of the V0 Version
func APIV0Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "v0")
	w.Write([]byte("v0.1"))
}

// SetConfig set the global config
func SetConfig(conf cfg.Config) {
	config = conf

	util.SetLogging(config.LogLevel, config.EnableSyslog, config.AppName)
}


// Commands is the main function of this package
func Commands() *mux.Router {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/versions", APIVersions).Methods("GET")
	rtr.HandleFunc("/api", APIVersions).Methods("GET")
	rtr.HandleFunc("/health", apiHealth).Methods("GET")
	rtr.HandleFunc("/api/feed/v0", APIV0GetFeed).Methods("POST")

	return rtr
}
