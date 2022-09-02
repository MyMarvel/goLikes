package main

import (
	"likes_handler/controllers"
	_ "likes_handler/docs"
	"likes_handler/routes"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Likes API for Accounts
// @version         1.0
// @description     Allows users to add likes to other users.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Petrov Nikita
// @contact.url    https://github.com/MyMarvel
// @contact.email  nikita.petrov.programmer@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1.0

func main() {
	s := NewSettings()
	factory := controllers.NewFactory()
	controllers.InitControllersFactory(factory)
	controllers.InitDatabase(s.DbHost, s.DbPort, s.DbPass)
	routes.InitControllersFactory(factory)
	r := routes.GenerateRoutes()
	swagger := r.Group("/swagger")
	{
		swagger.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	r.Run(":" + s.WebPort)
}
