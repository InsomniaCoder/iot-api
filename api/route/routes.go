package route

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/iot/api")
	{
		api.POST("sensors")
	}
	return r
}
