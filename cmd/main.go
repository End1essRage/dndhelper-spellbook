package main

import (
	"fmt"

	server "github.com/end1essrage/dndhelper-spellbook"
	handler "github.com/end1essrage/dndhelper-spellbook/pkg/handler"
)

func main() {
	/*TODO
	Добавить логгер
	Добавить считывание конфига
	*/
	handlers := handler.NewHandler()

	srv := new(server.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		fmt.Printf("error while running http server %s", err.Error())
	}
}
