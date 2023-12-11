package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	// Setting the logger to a file
	log.Out = os.Stdout
	// file, err := os.OpenFile("./btlog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	// 	log.Out = file
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// }
}

var log = logrus.New()

func main() {
	routes()
}
