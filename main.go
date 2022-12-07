package main

import (
	"github.com/abdghn/alpha-indo-soft-be-test/database"
	"github.com/abdghn/alpha-indo-soft-be-test/handler"
	"github.com/abdghn/alpha-indo-soft-be-test/repository"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")

	redis := database.ConnectRedis()
	db := database.ConnectDB()
	repo := repository.NewRepository(db, redis)
	h := handler.NewHandler(repo)
	router := gin.Default()

	router.GET("/articles", h.ViewArticles)
	router.POST("/articles", h.CreateArticle)
	router.Run(":6000")

}
