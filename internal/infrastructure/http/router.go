package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", ExampleHandler)
	}
}
