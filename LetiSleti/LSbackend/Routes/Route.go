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

	admin.Use(Token.AdminAuthMiddleware())
	user.Use(Token.UserAuthMiddleware())

	public.POST("/user/register", Controllers.RegisterUser)
	public.POST("/user/login", Controllers.LoginUser)
	public.GET("/user/getAll", Controllers.GetAllUsers)
	public.GET("/user/by-email/:email", Controllers.GetUserByEmail)
	public.DELETE("/user/:userId", Controllers.DeleteUser)

	public.POST("/ticket", Controllers.CreateTicket)
	user.GET("/ticket/getAll", Controllers.GetAllTickets)

	admin.POST("/flight/register", Controllers.RegisterFlight)
	public.GET("/flight/change-places-left/:flightId/:count", Controllers.ChangePlacesLeft)
	public.GET("/flight/getAll", Controllers.GetAllFlights)
	public.POST("/flight/search", Controllers.SearchFlights)
	admin.DELETE("/flight/:flightId", Controllers.DeleteFlight)

	r.Run()
	return nil
}
