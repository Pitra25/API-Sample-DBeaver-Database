package app

import (
	server "Educational-API-DBeaver-Sample-Database"
	"Educational-API-DBeaver-Sample-Database/internal/handler/http"
	"Educational-API-DBeaver-Sample-Database/internal/repository"
	service "Educational-API-DBeaver-Sample-Database/internal/servise"
	"Educational-API-DBeaver-Sample-Database/pkg/messages"
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Educational-API-DBeaver-Sample-Database
// @version 1.0
// @description API
// @host localhost:8000
// @BasePath /
// @in header
func Go() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(logrus.TraceLevel)

	if err := initConfig(); err != nil {
		logrus.Fatal("error initialization config: ", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal("error load .env variables: ", err.Error())
	}

	db, err := init_SQLiteDB()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	repo := repository.New(db)
	service := service.New(repo)
	handlerInit := http.NewHandler(service)
	handler := http.New(service, handlerInit)

	srv := startApp(handler)
	closeApp(srv, db)

}

func startApp(h *http.Handler) *server.New {
	srv := new(server.New)

	go func() {
		if err := srv.Run(
			viper.GetString("server.port"),
			h.InitRoutes(),
		); err != nil {
			logrus.Fatal("error run http server: ", err.Error())
		}
	}()

	logrus.Info("API SERVER STARTED")
	logrus.Infof(
		"URL: http://%s:%s/%s",
		viper.GetString("server.host"),
		viper.GetString("server.port"),
		viper.GetString("server.swagger_url"),
	)

	return srv
}

func closeApp(srv *server.New, db *sql.DB) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatal("error occurred on server shutting down: ", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Error("error occurred on db connection close: ", err.Error(), messages.Fatal)
	}

	logrus.Print("SERVER SHUTTING DOWN")
}
