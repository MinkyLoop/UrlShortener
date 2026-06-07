package api

import "github.com/gin-gonic/gin"

func SetupRouter(h *Handlers) *gin.Engine {
	r := gin.New()

	r.Use(LoggerMiddleware())
	r.Use(RecoveryMiddleware())

	api := r.Group("/api")
	{
		api.POST("/shorten", h.Shorten)
	}

	r.GET("/:shortCode", h.Redirect)

	return r
}
