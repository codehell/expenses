package tests

import (
	appConfig "expenses/config"
	"expenses/models"
	"expenses/routes"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"log"
	"time"
)

func clearUserTable(db *pg.DB) {
	_, err := db.Model((*models.User)(nil)).Where("true").Delete()
	if err != nil {
		log.Fatalln(err)
	}
}

func clearExpensesTable(db *pg.DB) {
	_, err := db.Model((*models.Expense)(nil)).Where("true").Delete()
	if err != nil {
		log.Fatalln(err)
	}
}

func getRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	e := gin.Default()
	r := routes.Router(e)
	return r
}

func makeTokenString(SigningAlgorithm string, email string) string {
	var config appConfig.Config
	var key = []byte(config.GetConfig().TokenKey)
	if SigningAlgorithm == "" {
		SigningAlgorithm = "HS256"
	}
	token := jwt.New(jwt.GetSigningMethod(SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["orig_iat"] = time.Now().Unix()
	var tokenString string
	tokenString, _ = token.SignedString(key)
	return tokenString
}
