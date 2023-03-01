package handler

import (
	"github.com/gin-gonic/gin"
	"template-server/api/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(CORSMiddleware())
	auth := router.Group("/template", h.identify)
	{
		auth.POST("/create", h.createTemplate)
		auth.GET("/", h.getAll)
		auth.GET("/:id", h.getTemplate)

		auth.PATCH("/:id", h.updateTemplate)
		auth.DELETE("/:id", h.deleteTemplate)
	}
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
