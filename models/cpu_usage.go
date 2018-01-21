package models

import (
	"fmt"
	"slurm_statistics/config"
)

type CpuUsage struct {
	Name     string  `json:"name"`
	CpuHours float64 `json:"cpu_hours"`
}

func CpuHourPerUser(from int64, to int64) ([]*CpuUsage, error) {
	prefix := config.Conf.Database.TablePrefix
	query := fmt.Sprintf(`SELECT
   a_t.user,
   SUM((CAST(SUBSTRING(SUBSTRING_INDEX(j_t.tres_alloc, ',', 1), 3) AS UNSIGNED) * (j_t.time_end - j_t.time_start) /
        (60 * 60))) AS cpu_hour
 FROM %s_assoc_table AS a_t
   JOIN %s_job_table AS j_t ON a_t.id_assoc = j_t.id_assoc
 WHERE j_t.time_start > ? AND j_t.time_end < ? AND time_end != 0 AND time_start != 0 AND
       j_t.tres_alloc IS NOT NULL AND j_t.tres_alloc != ''
 GROUP BY a_t.user
 ORDER BY cpu_hour DESC;`, prefix, prefix)
	return fetchCpuUsageByQuery(query, from, to)
}

func CpuHourPerAccount(from int64, to int64) ([]*CpuUsage, error) {
	prefix := config.Conf.Database.TablePrefix
	query := fmt.Sprintf(`SELECT
  a_t.acct,
  SUM((CAST(SUBSTRING(SUBSTRING_INDEX(j_t.tres_alloc, ',', 1), 3) AS UNSIGNED) * (j_t.time_end - j_t.time_start) /
       (60 * 60))) AS cpu_hour
FROM %s_assoc_table AS a_t
  JOIN %s_job_table AS j_t ON a_t.id_assoc = j_t.id_assoc
WHERE j_t.time_start > ? AND j_t.time_end < ? AND time_end != 0 AND time_start != 0 AND
      j_t.tres_alloc IS NOT NULL AND j_t.tres_alloc != ''
GROUP BY a_t.acct
ORDER BY cpu_hour DESC;`, prefix, prefix)
	return fetchCpuUsageByQuery(query, from, to)
}

func CpuHourPerPartition(from int64, to int64) ([]*CpuUsage, error) {
	prefix := config.Conf.Database.TablePrefix
	query := fmt.Sprintf(`SELECT
  j_t.partition,
  SUM((CAST(SUBSTRING(SUBSTRING_INDEX(j_t.tres_alloc, ',', 1), 3) AS UNSIGNED) * (j_t.time_end - j_t.time_start) /
       (60 * 60))) AS cpu_hour
FROM %s_job_table AS j_t
WHERE j_t.time_start > ? AND j_t.time_end < ? AND time_end != 0 AND time_start != 0 AND
      j_t.tres_alloc IS NOT NULL AND j_t.tres_alloc != ''
GROUP BY j_t.partition
ORDER BY cpu_hour DESC;`, prefix)
	return fetchCpuUsageByQuery(query, from, to)
}

func CpuHoursFull(from int64, to int64) (float64, error) {
	prefix := config.Conf.Database.TablePrefix
	query := fmt.Sprintf(`SELECT
  SUM((CAST(SUBSTRING(SUBSTRING_INDEX(j_t.tres_alloc, ',', 1), 3) AS UNSIGNED) * (j_t.time_end - j_t.time_start) /
       (60 * 60))) AS cpu_hour
FROM %s_job_table AS j_t
WHERE j_t.time_start > ? AND j_t.time_end < ? AND time_end != 0 AND time_start != 0 AND
      j_t.tres_alloc IS NOT NULL AND j_t.tres_alloc != '';`, prefix)
	stmt, err := db.Prepare(query)
	if err != nil {
		return -1, err
	}
	rows, err := stmt.Query(from, to)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	var result float64
	rows.Next()
	rows.Scan(&result)
	return result, nil
}

func fetchCpuUsageByQuery(query string, from int64, to int64) ([]*CpuUsage, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return deserializeCpuUsageRows(rows), nil
}
