package main

import (
	server "github.com/end1essrage/dndhelper-spellbook"
	handler "github.com/end1essrage/dndhelper-spellbook/pkg/handler"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("error while reading config %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error while reading environment %s", err.Error())
	}

	handlers := handler.NewHandler()

	srv := new(server.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error while reading environment %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
