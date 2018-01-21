package middleware

import (
	"github.com/gorilla/mux"
	"slurm_statistics/handler"
	"net/http"
	"fmt"
)

const DateTimeRegex = "\\d{4}[-]?\\d{1,2}[-]?\\d{1,2} \\d{1,2}:\\d{1,2}:\\d{1,2}"

/*
RegisterRoutes register routes and returns router
 */
func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.IndexHandler).Methods("GET")
	router.HandleFunc("/user_acct_stats", handler.UserAcctStatPage).Methods("GET")
	router.HandleFunc("/load", handler.LoadPageHandler).Methods("GET")
	router.HandleFunc("/resource_intensive_jobs", handler.ResourceIntensiveJobsPage).
		Methods("GET")

	router.HandleFunc("/api/user/jobs", handler.JobCountPerUserHandler).
		Queries(pairRegexps()).Methods("GET")
	router.HandleFunc("/api/account/jobs", handler.JobCountPerAcctHandler).
		Queries(pairRegexps()).Methods("GET")

	router.HandleFunc("/api/user/cpu", handler.CpuUsagePerUserHandler).
		Queries(pairRegexps()).Methods("GET")
	router.HandleFunc("/api/account/cpu", handler.CpuUsagePerAccountHandler).
		Queries(pairRegexps()).Methods("GET")

	router.HandleFunc("/api/partition/jobs", handler.LoadJobCount).
		Queries(pairRegexps()).Methods("GET")
	router.HandleFunc("/api/partition/cpu", handler.LoadCpuUsage).
		Queries(pairRegexps()).Methods("GET")
	router.HandleFunc("/api/job/intensive", handler.ResourceIntensiveJobs).Queries(
		"from", valueRegexp("from", DateTimeRegex), "to", valueRegexp("to", DateTimeRegex),
		"limit", valueRegexp("limit", "[0-9]+")).Methods("GET")


	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs)).Methods("GET")

	return router
}

func pairRegexps() (string, string, string, string) {
	return "from", valueRegexp("from", DateTimeRegex), "to", valueRegexp("to", DateTimeRegex)
}

func valueRegexp(key string, regexp string) string {
	return fmt.Sprintf("{%s:%s}", key, regexp)
}
