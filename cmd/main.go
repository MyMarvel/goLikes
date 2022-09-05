package main

import (
	"fmt"
	"likes_handler/configs"
	_ "likes_handler/docs"
	"likes_handler/internal/controllers"
	"likes_handler/internal/gRPCServer"
	"likes_handler/tools/proto"
	"likes_handler/web/app/routes"
	"net"
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	config "github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
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
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	if err := configs.Init(); err != nil {
		log.Fatal().Err(err)
	}
	factory := controllers.NewFactory()
	controllers.InitControllersFactory(factory)
	controllers.InitDatabase(config.GetString("database.host"), config.GetInt("database.port"), config.GetString("database.pass"))

	var wg sync.WaitGroup

	if config.GetBool("webAPI.enabled") {
		wg.Add(1)
		go func() {
			routes.InitControllersFactory(factory)
			r := routes.GenerateRoutes()
			swagger := r.Group("/swagger")
			{
				swagger.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
			}
			r.Run(":" + config.GetString("webAPI.port"))
		}()
	}

	if config.GetBool("gRPC.enabled") {
		wg.Add(1)
		go func() {
			log.Print("Starting gRPC server...")
			gRPCServer.InitControllersFactory(factory)
			s := grpc.NewServer()
			srv := &gRPCServer.GRPCServer{}
			proto.RegisterLikesServer(s, srv)

			l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.GetString("gRPC.host"), config.GetInt("gRPC.port")))
			if err != nil {
				log.Fatal().Err(err)
			}

			log.Printf("gRPC server started on port %d", config.GetInt("gRPC.port"))
			if err := s.Serve(l); err != nil {
				log.Fatal().Err(err)
			}
		}()
	}

	wg.Wait()
}
