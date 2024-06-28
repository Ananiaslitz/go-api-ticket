package main

import (
	"github.com/ananiaslitz/go-api-ticket/internal/infrastructure/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	http.RegisterRoutes(r)
	r.Run(":8080") // Inicia o servidor na porta 8080
}
