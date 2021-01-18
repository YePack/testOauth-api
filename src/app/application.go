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
	atServices := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atServices)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run("127.0.0.1:8080")
}
