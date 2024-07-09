package http

import (
	"net/http"

	"github.com/ananiaslitz/go-api-ticket/internal/domain/model"
	"github.com/gin-gonic/gin"
)

func ExampleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func GetTicketHandler(c *gin.Context) {
	// Criar prêmios
	prize1 := model.Prize{Name: "Gift Card", RequiredPoints: 100}
	prize2 := model.Prize{Name: "Free Coffee", RequiredPoints: 50}

	// Criar configuração
	config := model.Config{
		Prizes:    []model.Prize{prize1, prize2},
		MaxPoints: 100,
	}

	// Criar empresa
	company := model.Company{
		ID:   1,
		Name: "Awesome Company",
	}

	// Criar pontos
	points := model.Points{
		CurrentPoints: 75,
	}

	// Criar ticket
	ticket := model.Ticket{
		ID:      1,
		Points:  points,
		Company: company,
		Config:  config,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Controller de Ticket!",
		"data":    ticket,
	})
}
