package api

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/mmcdole/gofeed"
	"encoding/json"
	"io/ioutil"


	cfg "go-api-jsonrssfeed/types"
)

// APIV0GetFeed convert the given feed into json
func APIV0GetFeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error("Error:", err)
		return
	}

	var resGetFeed cfg.GetFeed
	err = json.Unmarshal(body, &resGetFeed)
	if err != nil {
		logrus.Error("Get URL: ", err)
		return
	}

	logrus.Debug(resGetFeed)

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(resGetFeed.Url)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "v0")

	w.Write([]byte(feed.String()))
}
