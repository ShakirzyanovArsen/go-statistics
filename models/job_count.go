package models

import (
	"fmt"
	"slurm_statistics/config"
)

type JobCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func JobCountPerAccount(from int64, to int64) ([]*JobCount, error) {
	prefix := config.Conf.Database.TablePrefix
	query := fmt.Sprintf(`SELECT
  a_t.acct,
  count(*) AS job_count
FROM %s_assoc_table AS a_t
  JOIN %s_job_table AS j_t ON a_t.id_assoc = j_t.id_assoc
WHERE j_t.time_start > ? AND j_t.time_end < ?
GROUP BY a_t.acct
ORDER BY job_count DESC;`, prefix, prefix)
	return fetchJobCountsByQuery(query, from, to)
}

func JobCountPerUser(from int64, to int64) ([]*JobCount, error) {
	prefix := config.Conf.Database.TablePrefix
	query := fmt.Sprintf(`SELECT
  a_t.user,
  count(*) AS job_count
FROM %s_assoc_table AS a_t
  JOIN %s_job_table AS j_t ON a_t.id_assoc = j_t.id_assoc
WHERE j_t.time_start > ? AND j_t.time_end < ?
GROUP BY a_t.user
ORDER BY job_count DESC;`, prefix, prefix)
	return fetchJobCountsByQuery(query, from, to)
}

func JobCountPerPartition(from int64, to int64) ([]*JobCount, error) {
	prefix := config.Conf.Database.TablePrefix
	query := fmt.Sprintf(`SELECT
  j_t.partition,
  count(*) AS job_count
FROM %s_job_table AS j_t
WHERE j_t.time_start > ? AND j_t.time_end < ?
GROUP BY j_t.partition
ORDER BY job_count DESC;`, prefix)
	return fetchJobCountsByQuery(query, from, to)
}

func JobCountFull(from int64, to int64) (int, error) {
	prefix := config.Conf.Database.TablePrefix
	query := fmt.Sprintf(`SELECT
  count(*) AS job_count
FROM %s_job_table AS j_t
WHERE j_t.time_start > ? AND j_t.time_end < ?;`, prefix)
	stmt, err := db.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(from, to)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	var result int
	rows.Next()
	rows.Scan(&result)
	return result, nil
}

func fetchJobCountsByQuery(query string, from int64, to int64)([]*JobCount, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return deserializeJobCountRows(rows), nil
}

