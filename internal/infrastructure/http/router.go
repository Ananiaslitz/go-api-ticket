package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", ExampleHandler)
		tickets := v1.Group("/tickets")
		{
			tickets.GET("", GetTicketHandler)
			v1.POST("", CreateTicket)
			v1.GET("/:id", GetTicket)
			v1.GET("", GetAllTickets)
			v1.PUT("/:id", UpdateTicket)
			v1.DELETE("/:id", DeleteTicket)
		}
	}
}
