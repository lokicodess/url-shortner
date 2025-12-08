package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lokicodess/url-shortner/internal/model"
)

type app struct {
	logger   *slog.Logger
	db       *sql.DB
	urlModel model.UrlModel
}

func main() {

	addr := flag.String("addr", ":8080", "Server address port")
	dsn := flag.String("dsn", "root:password@tcp(127.0.0.1:3306)/prod?parseTime=true", "Data Soure Name")

	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	db, err := openDB(*dsn)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("DB Started", "PORT", 3306)

	defer db.Close()
	app := &app{
		logger: logger,
		urlModel: model.UrlModel{
			DB: db,
		},
	}

	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("Server Listening", "PORT", *addr)

	err = srv.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
