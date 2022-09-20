package authserver

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	authorize "github.com/p1nant0m/gin-gin/controller/v1"
	"github.com/p1nant0m/gin-gin/log"
	"github.com/p1nant0m/gin-gin/pkg/middleware/auth"
	"github.com/p1nant0m/gin-gin/store/local"
)

func initRouter(g *gin.Engine) {
	initController(g)
}

func initController(g *gin.Engine) {
	authMiddleware := auth.NewJWTAuthMidleware()

	g.POST("/login", authMiddleware.LoginHandler)

	g.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Logger().For(c.Request.Context()).Infof("NoRoute claims: %#v\n", log.Warp(claims))
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	directIns, err := local.GetDirectLocalInsOr(local.GetLocalFactoryOrExit())
	if err != nil {
		log.Logger().Bg().Fatalf("get nil local instance\n", nil)
	}

	apiv1 := g.Group("/v1", authMiddleware.MiddlewareFunc())
	{
		// Refresh time can be longer than token timeout
		apiv1.GET("/refresh_token", authMiddleware.RefreshHandler)

		authzController := authorize.NewAuthzController(directIns)
		apiv1.POST("/authz", authzController.Authorize)
	}

}
