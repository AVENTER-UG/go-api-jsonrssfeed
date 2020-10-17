package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func apiHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	logrus.Debug("Health Check")

	w.WriteHeader(http.StatusOK)
}
