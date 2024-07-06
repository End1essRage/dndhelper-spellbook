package main

import (
	server "github.com/end1essrage/dndhelper-spellbook"
	client "github.com/end1essrage/dndhelper-spellbook/pkg/client"
	handler "github.com/end1essrage/dndhelper-spellbook/pkg/handler"
	"github.com/end1essrage/dndhelper-spellbook/pkg/service"
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

	//прокинуть из конфига
	client := client.NewClient("www.dnd5eapi.co")
	service := service.NewService(client)
	handlers := handler.NewHandler(service)

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
