package handler

import (
	"net/http"
	"slurm_statistics/tmpl"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ProcessTemplate(w, "index.html", nil)
}