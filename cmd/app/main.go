package main

import (
	"github.com/ananiaslitz/go-api-ticket/internal/infrastructure/http"
	"github.com/ananiaslitz/go-api-ticket/internal/infrastructure/persistence"
	"github.com/gin-gonic/gin"
)

func main() {
	db := persistence.InitDB()
	ticketRepo := persistence.NewTicketRepository(db)
	http.SetTicketRepository(ticketRepo)

	r := gin.Default()
	http.RegisterRoutes(r)
	r.Run(":8080")
}
