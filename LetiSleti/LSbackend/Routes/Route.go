package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Controllers"
)

// veliko slovo javna metoda
// malo slovo privatna metoda
func InitRoutes() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/user/register", Controllers.RegisterUser)

	r.Run()
	return nil
}
