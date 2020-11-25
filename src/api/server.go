package api

import (
	"gorestapimysql/src/config"
	"gorestapimysql/src/handler"

	"github.com/gin-gonic/gin"
)

//Run function invocked to start application
func Run() {
	// gin.ForceConsoleColor()
	router := gin.New()
	config.SetDB()
	// use custom logger
	router.Use(gin.LoggerWithFormatter(handler.Log))
	router.Use(gin.Recovery())
	router.GET("/transactions", handler.TransactionGet)
	router.Run()
}