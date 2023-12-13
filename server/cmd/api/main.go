package main

import (
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type application struct {
	config config
	logger logrus.Logger
}

type config struct {
	port int
	env  string
}

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
	var cfg config
	cfg.port = 4000

	app := &application{
		logger: *logrus.StandardLogger(),
		config: cfg,
	}

	srv := &http.Server{
		Handler:      app.routes(),
		Addr:         "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	app.logger.Infof("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	app.logger.Error(err)

}
