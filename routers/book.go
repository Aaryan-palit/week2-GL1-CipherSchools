package routers

import (
	"github.com/DigantaChauduri06/week2-GL1-CipherSchools/database"
	"github.com/DigantaChauduri06/week2-GL1-CipherSchools/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default() // get the dfault engine for further customization
	api := handler.Handler{
		DB: database.GetDB(), //set the handler DB
	}
	router.GET("/books", api.GetBooks) //set thefunction for thiss  https://localhost:8080/books : Get Request
	// while calling the handler function , gin will pass *gin.Context in the handler function
	router.POST("/books", api.SaveBook)

	return router

}
