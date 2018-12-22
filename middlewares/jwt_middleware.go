package middlewares

import (
	appConfig "expenses/config"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"

	"expenses/models"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

var identityKey = "email"

func JwtMiddleware() (*jwt.GinJWTMiddleware, error) {
	var config appConfig.Config
	confFile := "./config.yaml"
	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal(err)
	}
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(config.TokenKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			email, password, ok := c.Request.BasicAuth()
			if !ok {
				return "", jwt.ErrMissingLoginValues
			}
			user := models.User{Email: email}
			err := user.GetUserByEmail()
			if err != nil {
				return "", err
			}
			hashed := []byte(user.Password)
			err = bcrypt.CompareHashAndPassword(hashed, []byte(password))
			if err != nil {
				return "", jwt.ErrFailedAuthentication
			} else {
				return &models.User{
					Email: email,
				}, nil
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*models.User); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.Status(code)
			// c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware, nil
}
