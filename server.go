package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"slurm_statistics/config"
	"slurm_statistics/logging"
	"slurm_statistics/middleware"
	"slurm_statistics/utils"
)

func main() {
	r := middleware.RegisterRoutes()
	config.Init()
	logging.Init()
	address := fmt.Sprintf(":%d", config.Conf.Port)
	log.WithField("address", address).Info("starting server")
	err := http.ListenAndServe(address, r)
	utils.CheckFatal(err)
}
