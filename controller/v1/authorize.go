package authorize

import (
	"errors"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/p1nant0m/gin-gin/authorization"
	"github.com/p1nant0m/gin-gin/authorization/authorizer"
	userv1 "github.com/p1nant0m/gin-gin/pkg/api/v1"
	"github.com/p1nant0m/gin-gin/pkg/core"
	"github.com/sirupsen/logrus"
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
		logrus.Warning("expected userv1.User struct, got unknown type")
		core.WriteResponse(c, errors.New("unknown type received"), nil)
	} else {
		auth := authorization.NewAuthorizer(authorizer.NewAuthorization(a.datastore))
		resp := auth.Authorize(u)
		core.WriteResponse(c, nil, resp)
	}
}
