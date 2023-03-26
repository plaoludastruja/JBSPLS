package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Controllers"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Helper/Cors"
	"github.com/plaoludastruja/JBSPLS/LetiSleti/LSbackend/Helper/Token"
)

// veliko slovo javna metoda
// malo slovo privatna metoda
func InitRoutes() *gin.Engine {

	r := gin.New()
	r.Use(Cors.CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	public := r.Group("")
	admin := r.Group("admin")
	user := r.Group("user")
	//TODO
	admin.Use(Token.AdminAuthMiddleware())
	user.Use(Token.UserAuthMiddleware())

	public.POST("/user/register", Controllers.RegisterUser)
	public.POST("/user/login", Controllers.LoginUser)

	public.POST("/ticket", Controllers.CreateTicket)
	public.GET("/ticket/getAll", Controllers.GetAllTickets)

	public.POST("/flight/register", Controllers.RegisterFlight)
	public.GET("/flight/change-places-left/:flightId", Controllers.ChangePlacesLeft)

	public.GET("/user/getAll", Controllers.GetAllUsers)
	public.GET("/user/by-email/:email", Controllers.GetUserByEmail)
	public.GET("/flight/getAll", Controllers.GetAllFlights)

	public.DELETE("/user/:userId", Controllers.DeleteUser)
	r.Run()
	return nil
}
