package config

import (
	"github.com/tkanos/gonfig"
	"slurm_statistics/utils"
)

type Configuration struct {
	Port int
}

var Conf *Configuration

func Init() {
	Conf = &Configuration{}
	err := gonfig.GetConf("config/config.json", Conf)
	utils.CheckFatal(err)
}
