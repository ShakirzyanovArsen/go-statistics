package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"slurm_statistics/config"
	"slurm_statistics/logging"
	"slurm_statistics/middleware"
	"slurm_statistics/utils"
	"slurm_statistics/models"
)

func main() {
	config.Init()
	logging.Init()
	initDb()
	address := fmt.Sprintf(":%d", config.Conf.Port)
	log.WithField("address", address).Info("starting server")
	r := middleware.RegisterRoutes()
	err := http.ListenAndServe(address, r)
	utils.CheckFatal(err)
}

func initDb() {
	dbConf := config.Conf.Database;
	dataSourceName := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?autocommit=true`, dbConf.User, dbConf.Password,
		dbConf.Host, dbConf.Port, dbConf.Database)
	models.InitDB(dataSourceName)
}
