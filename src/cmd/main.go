package main

import (
	application "AntonyChR/golang-hexagonal-arch-template/internal/core/application"
	httpServer "AntonyChR/golang-hexagonal-arch-template/internal/core/infrastructure/http"
	repositorie "AntonyChR/golang-hexagonal-arch-template/internal/core/infrastructure/repositories"
)

func main() {

	// dbConfig := repositorie.MysqlConfig{
	// 	Database: "test",
	// 	Table:    "todos",
	// 	Host:     "192.168.18.46",
	// 	Port:     "3306",
	// 	User:     "root",
	// }

	todoService := application.TodoService{
		//Repo: repositorie.InitializeMysqlRepository(dbConfig),
		Repo: repositorie.InitializeMemoryRepository("todos", "test"),
	}

	server := httpServer.Initialize(":3000", &todoService)
	server.Start()
}
