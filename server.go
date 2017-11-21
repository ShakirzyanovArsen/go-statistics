package main

import (
	"slurm_statistics/config"
	"slurm_statistics/middleware"
	"slurm_statistics/utils"
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
	"slurm_statistics/logging"
)

func main() {
	r := middleware.RegisterRoutes()
	config.Init()
	logging.Init()
	address := fmt.Sprintf(":%d", config.Conf.Port)
	log.WithField("address", address).Info("starting server")
	utils.CheckFatal(http.ListenAndServe(address, r))
}
