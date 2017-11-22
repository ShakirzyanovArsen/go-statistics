package config

import (
	"github.com/tkanos/gonfig"
	"slurm_statistics/utils"
)

//Configuration structure for config
type Configuration struct {
	Port int
}

//Conf - pointer to configuration
var Conf *Configuration

//Init loads config from config.json and save result in Conf var
func Init() {
	Conf = &Configuration{}
	err := gonfig.GetConf("config/config.json", Conf)
	utils.CheckFatal(err)
}
