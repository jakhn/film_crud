package api

import (
	_ "github.com/jakhn/film_crud/api/docs"
	"github.com/jakhn/film_crud/api/handler"
	"github.com/jakhn/film_crud/config"
	"github.com/jakhn/film_crud/pkg/helper"
	"github.com/jakhn/film_crud/storage"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpApi(cfg *config.Config, r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandlerV1(cfg ,storage)

	r.Use(customCORSMiddleware())

	v1 := r.Group("/v1")
	r.POST("/login", handlerV1.Login)
	r.POST("/loginadmin", handlerV1.LoginAdmin) 

	v1.Use(check()) 
	r.POST("/book", handlerV1.CreateBook)
	r.GET("/book/:id", handlerV1.GetBookById)
	r.GET("/book", handlerV1.GetBookList)
	r.PUT("/book/:id", handlerV1.UpdateBook)
	r.DELETE("/book/:id", handlerV1.DeleteBook)

	r.POST("/user", handlerV1.CreateUser)
	r.GET("/user/:id", handlerV1.GetUserById)
	v1.GET("/user", handlerV1.GetUserList)
	r.PUT("/user/:id", handlerV1.UpdateUser)
	r.DELETE("/user/:id", handlerV1.DeleteUser)

	v1.POST("/order", handlerV1.CreateOrder)
	v1.GET("/order/:id", handlerV1.GetOrderById)
	v1.GET("/order", handlerV1.GetOrderList)
	v1.PUT("/order/:id", handlerV1.UpdateOrder)
	v1.DELETE("/order/:id", handlerV1.DeleteOrder)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func check() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, ok := ctx.Request.Header["Authorization"]; ok {
			_, err := helper.ExtractClaims(ctx.Request.Header["Authorization"][0], config.Load().AuthSecretKey)
			_, err2 := helper.ExtractClaims(ctx.Request.Header["Authorization"][0], config.Load().SuperAdmin)
			if err != nil && err2 != nil {
				ctx.AbortWithError(http.StatusForbidden, errors.New("not found password"))
				return
			} else {
				ctx.Next()
			}
		}
	}
}



func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
