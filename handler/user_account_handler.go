package handler

import (
	"net/http"
	"slurm_statistics/utils"
	"slurm_statistics/models"
	log "github.com/sirupsen/logrus"
	"slurm_statistics/tmpl"
)

func JobCountPerUserHandler(w http.ResponseWriter, r *http.Request) {
	from, to := utils.MakeTimesFromQuery(r)
	jobCount, err := models.JobCountPerUser(from.Unix(), to.Unix())
	utils.CheckError(err)
	utils.WriteAsJson(w, jobCount)
}

func JobCountPerAcctHandler(w http.ResponseWriter, r *http.Request) {
	from, to := utils.MakeTimesFromQuery(r)
	jobCount, err := models.JobCountPerAccount(from.Unix(), to.Unix())
	utils.CheckError(err)
	utils.WriteAsJson(w, jobCount)
}

func CpuUsagePerUserHandler(w http.ResponseWriter, r *http.Request) {
	from, to := utils.MakeTimesFromQuery(r)
	cpuUsage, err := models.CpuHourPerUser(from.Unix(), to.Unix())
	utils.CheckError(err)
	utils.WriteAsJson(w, cpuUsage)
}

func CpuUsagePerAccountHandler(w http.ResponseWriter, r *http.Request) {
	from, to :=utils.MakeTimesFromQuery(r)
	cpuUsage, err := models.CpuHourPerAccount(from.Unix(), to.Unix())
	utils.CheckError(err)
	utils.WriteAsJson(w, cpuUsage)
}

func UserAcctStatPage(w http.ResponseWriter, r *http.Request) {
	log.Info("UserAcctStatPage")
	tmpl.ProcessTemplate(w, "user_account_stat.html", nil)
}


