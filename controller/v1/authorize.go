package authorize

import (
	"errors"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/p1nant0m/gin-gin/authorization"
	"github.com/p1nant0m/gin-gin/authorization/authorizer"
	"github.com/p1nant0m/gin-gin/log"
	userv1 "github.com/p1nant0m/gin-gin/pkg/api/v1"
	"github.com/p1nant0m/gin-gin/pkg/core"
	"go.opentelemetry.io/otel/attribute"
)

type AuthzController struct {
	datastore authorizer.PolicyGetter
}

func NewAuthzController(datastore authorizer.PolicyGetter) *AuthzController {
	return &AuthzController{datastore}
}

func (a *AuthzController) Authorize(c *gin.Context) {
	user, _ := c.Get(jwt.IdentityKey)
	if u, ok := user.(*userv1.User); !ok {
		log.Logger().For(c.Request.Context()).Warningf("expected userv1.User struct, got unknown type\n", nil, attribute.String("controller.name", "authorize"))
		core.WriteResponse(c, errors.New("unknown type received"), nil)
	} else {
		auth := authorization.NewAuthorizer(authorizer.NewAuthorization(a.datastore))
		resp := auth.Authorize(u)
		log.Logger().For(c.Request.Context()).Infof("user %v is authenticated. \n", log.Warp(u.UserName),
			attribute.String("autz.resp", resp.ToString()))
		core.WriteResponse(c, nil, resp)
	}
}
