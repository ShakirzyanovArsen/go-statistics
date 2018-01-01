package handler

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"slurm_statistics/tmpl"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Index request")
	tmpl.ProcessTemplate(w, "index.html")
}