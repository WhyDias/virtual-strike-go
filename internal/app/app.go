package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"virtual-strike-backend-go/internal/handler"
	"virtual-strike-backend-go/internal/server"
	"virtual-strike-backend-go/internal/service"
	"virtual-strike-backend-go/pkg/models"
	monitoring "virtual-strike-backend-go/pkg/moniroting"
)

func Run() {

	monitoring.Init()

	models.ConnectDataBase()

	if err := initConfig(); err != nil {
		logrus.Fatalf("error occured while initializing configs: %s", err.Error())
	}

	err := viper.ReadInConfig()

	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	services := service.NewService()
	handlers := handler.NewHandler(services)
	server := new(server.Server)

	go func() {
		if err := server.Run(viper.GetString("PORT"), handlers.InitRoutes()); err != nil {
			logrus.Errorf("error occured while running http server %s/n", err.Error())
		}
	}()

	logrus.Printf("server is starting at port: %s", viper.GetString("PORT"))

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	logrus.Printf("Server is stopping at port: %s", "8877")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("./configs")
	env := os.Getenv("ENV")
	if env == "production" {
		viper.SetConfigName("configs")
		logrus.Info("loaded production configuration")
	} else {
		viper.SetConfigName("devconfig")
		logrus.Info("loaded dev configuration")
	}
	return viper.ReadInConfig()
}
