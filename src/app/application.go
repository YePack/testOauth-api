package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yepack/testOauth-api/src/domain/access_token"
	"github.com/yepack/testOauth-api/src/http"
	"github.com/yepack/testOauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {

	atHandler := http.NewAccessTokenHandler(access_token.NewService(db.NewRepository()))
	router.POST("/oauth/access_token", atHandler.Create)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run("127.0.0.1:8080")

}
