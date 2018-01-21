package models

import (
	"slurm_statistics/config"
	"fmt"
)

type Job struct {
	Inx         int     `json:"inx"`
	Account     string  `json:"account"`
	Partition   string  `json:"partition"`
	User        string  `json:"user"`
	TimeStart   string  `json:"time_start"`
	TimeEnd     string  `json:"time_end"`
	JobTime     string  `json:"job_time"`
	CpuHours    float64 `json:"cpu_hours"`
	CpuRequired int     `json:"cpu_required"`
}

func ResourceIntensiveJobs(from int64, to int64, limit int) ([]*Job, error) {
	prefix := config.Conf.Database.TablePrefix
	dateTimeFormat := "%Y-%m-%d %H:%i:%s"
	timeIntervalFormat := "%Hh %im"
	query := fmt.Sprintf(`SELECT j_t.job_db_inx, j_t.partition, j_t.account, a_t.user,
  FROM_UNIXTIME(j_t.time_start, '%s') AS time_start,
  FROM_UNIXTIME(j_t.time_end, '%s') AS time_end,
  TIME_FORMAT(SEC_TO_TIME(j_t.time_end - j_t.time_start),'%s') AS job_time,
  (CAST(SUBSTRING(SUBSTRING_INDEX(j_t.tres_alloc, ',', 1), 3) AS UNSIGNED) * (j_t.time_end - j_t.time_start) /
       (60 * 60)) AS cpu_hour,
  SUBSTRING(SUBSTRING_INDEX(j_t.tres_alloc, ',', 1), 3) AS cpu_requred
FROM %s_job_table AS j_t
  JOIN %s_assoc_table AS a_t ON a_t.id_assoc = j_t.id_assoc
WHERE j_t.time_start > ? AND j_t.time_end < ? AND time_end != 0 AND time_start != 0 AND
      j_t.tres_alloc IS NOT NULL AND j_t.tres_alloc != ''
ORDER BY cpu_hour DESC LIMIT ?;`, dateTimeFormat, dateTimeFormat, timeIntervalFormat, prefix, prefix)
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(from, to, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	jobs := make([]*Job, 0)
	for rows.Next() {
		job := new(Job)
		rows.Scan(&job.Inx, &job.Partition, &job.Account, &job.User, &job.TimeStart, &job.TimeEnd,
			&job.JobTime, &job.CpuHours, &job.CpuRequired)
		jobs = append(jobs, job)
	}
	return jobs, nil
}
