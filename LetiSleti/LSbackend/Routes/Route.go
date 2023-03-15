package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Controllers"
)

// veliko slovo javna metoda
// malo slovo privatna metoda
func InitRoutes() *gin.Engine {

	r := gin.New()
	r.Use(CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/user/register", Controllers.RegisterUser)
	r.GET("/user/getAll", Controllers.GetAllUsers)

	r.Run()
	return nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
