package authserver

import (
	"net/http"
	"os"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	authorize "github.com/p1nant0m/gin-gin/controller/v1"
	"github.com/p1nant0m/gin-gin/pkg/middleware/auth"
	"github.com/p1nant0m/gin-gin/store/local"
	"github.com/sirupsen/logrus"
)

func Run() {
	port := os.Getenv("PORT")
	r := gin.Default()

	if port == "" {
		port = "8000"
	}

	logrus.SetLevel(logrus.DebugLevel)

	authMiddleware := auth.NewJWTAuthMidleware()

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		logrus.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	directIns, err := local.GetDirectLocalInsOr(local.GetLocalFactoryOrExit())
	if err != nil {
		logrus.Fatalln("get nil local instance")
	}

	apiv1 := r.Group("/v1", authMiddleware.MiddlewareFunc())
	{
		// Refresh time can be longer than token timeout
		apiv1.GET("/refresh_token", authMiddleware.RefreshHandler)

		authzController := authorize.NewAuthzController(directIns)
		apiv1.POST("/authz", authzController.Authorize)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		logrus.Fatal(err)
	}
}
