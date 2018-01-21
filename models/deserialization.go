package models

import "database/sql"

func deserializeCpuUsageRows(rows *sql.Rows) []*CpuUsage {
	cpuUsageSlice := make([]*CpuUsage, 0)
	for rows.Next(){
		cpuUsage := new(CpuUsage)
		rows.Scan(&cpuUsage.Name, &cpuUsage.CpuHours)
		cpuUsageSlice = append(cpuUsageSlice, cpuUsage)
	}
	return cpuUsageSlice
}

func deserializeJobCountRows(rows *sql.Rows) []*JobCount {
	jobCountSlice := make([]*JobCount, 0)
	for rows.Next() {
		jobCount := new(JobCount)
		rows.Scan(&jobCount.Name, &jobCount.Count)
		jobCountSlice = append(jobCountSlice, jobCount)
	}
	return jobCountSlice
}
