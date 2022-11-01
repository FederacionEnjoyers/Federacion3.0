package handler

import (
	"github.com/Serelllka/Federacion/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	corsConfig := cors.Config{
		AllowOrigins:     viper.GetStringSlice("cors.allowOrigins"),
		AllowMethods:     viper.GetStringSlice("cors.allowMethods"),
		AllowHeaders:     viper.GetStringSlice("cors.allowHeaders"),
		ExposeHeaders:    viper.GetStringSlice("cors.exposeHeaders"),
		AllowCredentials: viper.GetBool("cors.allowCredentials"),
		MaxAge:           viper.GetDuration("cors.maxAge"),
	}

	router.Use(cors.New(corsConfig))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		friends := api.Group("/friends")
		{
			friends.POST("/add-friend", h.AddFriend)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
