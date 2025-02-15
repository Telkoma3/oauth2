package utils

import (
    "os"

    log "github.com/sirupsen/logrus"
)

func InitLogger() {
    log.SetFormatter(&log.JSONFormatter{})
    log.SetOutput(os.Stdout)
    log.SetLevel(log.InfoLevel)
}

func Logger() *log.Logger {
    return log.StandardLogger()
}