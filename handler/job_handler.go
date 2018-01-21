package handler

import (
	"net/http"
	"slurm_statistics/utils"
	"slurm_statistics/models"
	"strconv"
	"slurm_statistics/tmpl"
)

func ResourceIntensiveJobsPage(w http.ResponseWriter, r *http.Request) {
	tmpl.ProcessTemplate(w, "resource_intensive_jobs.html", nil)
}

func ResourceIntensiveJobs(w http.ResponseWriter, r *http.Request) {
	from, to := utils.MakeTimesFromQuery(r)
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	utils.CheckError(err)
	jobs, err := models.ResourceIntensiveJobs(from.Unix(), to.Unix(), limit)
	utils.CheckError(err)
	utils.WriteAsJson(w, jobs)
}
