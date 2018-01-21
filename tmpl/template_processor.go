package tmpl

import (
	"slurm_statistics/utils"
	"html/template"
	"net/http"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func ProcessTemplate(w http.ResponseWriter, templateName string, variable interface{}) {
	w.Header().Set("Content-Type", "text/html")
	dir, err := os.Getwd()
	targetTemplate := fmt.Sprintf(`%s/tmpl/%s`, dir, templateName)
	log.Debug(targetTemplate)
	t, err := template.ParseFiles(targetTemplate, "tmpl/head.html")
	utils.CheckError(err)
	err = t.Execute(w, variable)
	utils.CheckError(err)
}
