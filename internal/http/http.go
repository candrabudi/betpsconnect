package http

import (
	District "betpsconnect/internal/app/district"
	Resident "betpsconnect/internal/app/resident"
	SubDistrict "betpsconnect/internal/app/subdistrict"
	User "betpsconnect/internal/app/user"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Here we define route function for user Handlers that accepts gin.Engine and factory parameters
func NewHttp(g *gin.Engine, f *factory.Factory) {

	Index(g)
	// Here we use logger middleware before the actual API to catch any api call from clients
	g.Use(gin.Logger())
	// Here we use the recovery middleware to catch a panic, if panic occurs recover the application witohut shutting it off
	g.Use(gin.Recovery())

	g.Use(middleware.CORSMiddleware())

	// Here we define a router group
	v1 := g.Group("/api/v1")
	Resident.NewHandler(f).Router(v1.Group("/resident"))
	District.NewHandler(f).Router(v1.Group("/district"))
	SubDistrict.NewHandler(f).Router(v1.Group("/subdistrict"))
	User.NewHandler(f).Router(v1.Group("/user"))

}

func Index(g *gin.Engine) {
	g.GET("/", func(context *gin.Context) {
		context.JSON(200, struct {
			Name string `json:"name"`
		}{
			Name: "Tps Connect",
		})
	})
}
