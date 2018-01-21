package handler

import (
	"net/http"
	"slurm_statistics/tmpl"
	"slurm_statistics/utils"
	"slurm_statistics/models"
)

type PartitionPage struct {
	Partitions []*string
}

func LoadPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ProcessTemplate(w, "load.html", nil)
}

func LoadJobCount(w http.ResponseWriter, r *http.Request) {
	from, to :=utils.MakeTimesFromQuery(r)
	jobCount, err := models.JobCountPerPartition(from.Unix(), to.Unix())
	utils.CheckError(err)
	fullJobCount, err := models.JobCountFull(from.Unix(), to.Unix())
	fullJob := new(models.JobCount)
	fullJob.Name = "All partitions"
	fullJob.Count = fullJobCount
	result := make([]*models.JobCount, 0)
	result = append(result, fullJob)
	for _, cnt := range jobCount {
		result = append(result, cnt)
	}
	utils.WriteAsJson(w, result)
}

func LoadCpuUsage(w http.ResponseWriter, r *http.Request) {
	from, to :=utils.MakeTimesFromQuery(r)
	jobCount, err := models.CpuHourPerPartition(from.Unix(), to.Unix())
	utils.CheckError(err)
	fullCpuUsage, err := models.CpuHoursFull(from.Unix(), to.Unix())
	fullCpu := new(models.CpuUsage)
	fullCpu.Name = "All partitions"
	fullCpu.CpuHours = fullCpuUsage
	result := make([]*models.CpuUsage, 0)
	result = append(result, fullCpu)
	for _, cnt := range jobCount {
		result = append(result, cnt)
	}
	utils.WriteAsJson(w, result)
}
