package logging

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
}
