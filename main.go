package main

import (
	"os"

	"github.com/apsdehal/go-logger"
	"github.com/pkg/errors"
)

func main() {
	log, err := logger.New("test", 1, os.Stdout)
	if err != nil {
		panic(err)
	}

	err = errors.New("server error")

	log.Critical("Critical")
	log.Debug("Debug")
	log.Error("Error")
	log.Notice("Notice")
	log.Info("Info")
}
