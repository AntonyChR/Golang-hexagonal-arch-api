package main

import (
	controllers "chat/delivery/controllers"
	database "chat/delivery/database"
	middlewares "chat/delivery/middlewares"
	services "chat/domain/services"
	repositoryImplementation "chat/infrastructure/repositories"
	lib "chat/lib"
	ws "chat/ws"
	http "net/http"
	"os"

	"fmt"

	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
)

func main() {

	config, err := lib.ReadConfig("config.toml")
	CheckCriticalError(err)

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.AllowedOrigins
	corsConfig.AllowMethods = config.AllowedMethods
	corsConfig.AllowCredentials = config.AllowCredentials

	router.Use(cors.New(corsConfig))

	dbConfig := database.DBConfig{
		Database: config.DBName,
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
	}

	db, err := database.ConnectToDatabase(dbConfig)
	CheckCriticalError(err)

	userRepository := repositoryImplementation.NewMySqlUserRepository("users", db)
	messageRepository := repositoryImplementation.NewMySqlMessageRepository("messages", db)

	userService := services.NewUserService(userRepository)
	messageService := services.NewMessageService(messageRepository)

	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(userService)
	messageController := controllers.NewMessageController(messageService)

	router.LoadHTMLFiles("./public/index.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	swStore := ws.NewWesocketStore()
	go swStore.Run()

	protected := router.Group("/api")
	protected.Use(middlewares.AuhtRequired())
	{
		protected.GET("/ws", func(ctx *gin.Context) {
			ws.Server(swStore, ctx, messageService)
		})
		protected.GET("/user/all", userController.GetAllUsers)
		protected.GET("/messages", messageController.GetChat)
	}
	//router.GET("api/ws", func(ctx *gin.Context) {
	//		ws.Server(swStore, ctx, messageService)
	//})

	router.POST("api/user/", userController.CreateUser)
	router.DELETE("api/user/:id", userController.DeleteUserById)
	router.GET("api/user/:id", userController.GetUserById)

	router.POST("api/auth/login", authController.Login)
	router.GET("api/auth/revalidate", authController.RevalidateJWT)

	fmt.Println("Open: http://localhost" + config.ServerPort)

	router.Run(config.ServerPort)
}

func CheckCriticalError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
