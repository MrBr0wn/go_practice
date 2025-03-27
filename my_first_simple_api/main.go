package main

import (
	"github.com/gin-gonic/gin"
	"my_first_simple_api/storage"
	"my_first_simple_api/handler"
)

func main() {
	memoryStorage := storage.NewMemoryStorage()
	handler := handler.NewHandler(memoryStorage)

	router := gin.Default()

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run(":3001")
}
