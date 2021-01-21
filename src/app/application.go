package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yepack/testOauth-api/src/http"
	"github.com/yepack/testOauth-api/src/repository/db"
	"github.com/yepack/testOauth-api/src/repository/rest"
	"github.com/yepack/testOauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {

	atHandler := http.NewAccessTokenHandler(
		access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.POST("/oauth/access_token", atHandler.Create)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run("127.0.0.1:8080")

}
