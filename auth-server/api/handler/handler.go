package handler

import (
	"auth-server/api/service"
	"github.com/gin-gonic/gin"
)

type request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitHandler() *gin.Engine {
	router := gin.New()

	//resp.AddHeader(restful.HEADER_AccessControlAllowOrigin, "*")
	//chain.ProcessFilter(req, resp)

	router.Use(CORSMiddleware())
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
