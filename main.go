package main

import (
	"likes_handler/controllers"
	_ "likes_handler/docs"
	"likes_handler/routes"
	"likes_handler/settings"

	"log"

	config "github.com/spf13/viper"
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
	if err := settings.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	factory := controllers.NewFactory()
	controllers.InitControllersFactory(factory)
	controllers.InitDatabase(config.GetString("database.host"), config.GetString("database.port"), config.GetString("database.pass"))
	if config.GetBool("services.webAPI") {
		routes.InitControllersFactory(factory)
		r := routes.GenerateRoutes()
		swagger := r.Group("/swagger")
		{
			swagger.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		}
		r.Run(":" + config.GetString("webAPI.port"))
	}
	if config.GetBool("services.gRPC") {

	}
}
