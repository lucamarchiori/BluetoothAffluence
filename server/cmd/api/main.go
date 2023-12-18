package main

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"lucamarchiori.bluetoothAffluence/server/internal/data"
)

type application struct {
	config config
	logger logrus.Logger
	models data.Models
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
	var err error
	var statement *sql.Stmt

	cfg.port = 4000
	db, err := sql.Open("sqlite3", "../../database/ServerDB.db")

	if err != nil {
		log.Error(err)
		panic(1)
	}

	statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS devices (id INTEGER PRIMARY KEY, address TEXT, alias TEXT, name TEXT, txPower REAL, rssi REAL)")
	statement.Exec()
	statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS scanners (id INTEGER PRIMARY KEY, address TEXT, alias TEXT, name TEXT)")
	statement.Exec()

	app := &application{
		logger: *logrus.StandardLogger(),
		config: cfg,
		models: data.NewModels(db),
	}

	srv := &http.Server{
		Handler:      app.routes(),
		Addr:         "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	app.logger.Infof("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	app.logger.Error(err)

}
