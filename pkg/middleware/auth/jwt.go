package auth

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	userv1 "github.com/p1nant0m/gin-gin/pkg/api/v1"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var (
	KeyGenerate = func() []byte {
		key := make([]byte, 128)
		rand.Read(key)
		return key
	}
	emptyLoginInfo = &login{}
)

type PayloadFunc func(interface{}) jwt.MapClaims
type AuthorizatorFunc func(interface{}, *gin.Context) bool
type AuthenticatorFunc func(*gin.Context) (interface{}, error)

func NewJWTAuthMidleware() *jwt.GinJWTMiddleware {
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Gin-gin",
		Key:         KeyGenerate(),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: jwt.IdentityKey,
		PayloadFunc: getPayloadFunc(),
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &userv1.User{
				UserName: claims[jwt.IdentityKey].(string),
			}
		},
		Authenticator: getAuthenticator(),
		Authorizator:  getAuthorizator(),
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	return authMiddleware
}

func getPayloadFunc() PayloadFunc {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*userv1.User); ok {
			return jwt.MapClaims{
				jwt.IdentityKey: v.UserName,
			}
		}
		return jwt.MapClaims{}
	}
}

func getAuthorizator() AuthorizatorFunc {
	return func(data interface{}, c *gin.Context) bool {
		if _, ok := data.(*userv1.User); ok {
			return true
		}

		return false
	}
}

func parseHeader(c *gin.Context) (*login, error) {

	// Authorization Header is setting
	authInfo := strings.SplitN(c.GetHeader("Authorization"), " ", 2)
	if len(authInfo) < 2 {
		fmt.Printf("ClientIP: %v with authorization header but length less than 2\n", c.ClientIP())
		return emptyLoginInfo, jwt.ErrInvalidAuthHeader
	}

	if authInfo[0] != "Basic" {
		return emptyLoginInfo, jwt.ErrInvalidAuthHeader
	}

	payload, err := base64.StdEncoding.DecodeString(authInfo[1])
	if err != nil {
		fmt.Printf("ClientIP: %v base64 decode failuer err:%v \n", c.ClientIP(), err)
		return emptyLoginInfo, jwt.ErrInvalidAuthHeader
	}

	authInfo = strings.SplitN(string(payload), ":", 2)
	if len(authInfo) < 2 {
		fmt.Printf("ClientIP: %v invalid username and password\n", c.ClientIP())
		return emptyLoginInfo, jwt.ErrInvalidAuthHeader
	}

	// Format Validation success
	return &login{Username: authInfo[0], Password: authInfo[1]}, nil

}

func parseBody(c *gin.Context) (*login, error) {
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return emptyLoginInfo, jwt.ErrMissingLoginValues
	}

	return &loginVals, nil
}

func getAuthenticator() AuthenticatorFunc {
	return func(c *gin.Context) (interface{}, error) {
		var (
			loginInfo *login
			err       error
		)

		if len(c.Request.Header.Get("Authorization")) != 0 {
			loginInfo, err = parseHeader(c)
		} else {
			loginInfo, err = parseBody(c)
		}

		if err != nil {
			return nil, err
		}

		if (loginInfo.Username == "admin" && loginInfo.Password == "admin") || (loginInfo.Username == "test" && loginInfo.Password == "test") {
			return &userv1.User{
				UserName:  loginInfo.Username,
				LastName:  "Bo-Yi",
				FirstName: "Wu",
			}, nil
		}

		return nil, jwt.ErrFailedAuthentication
	}
}
