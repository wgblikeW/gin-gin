package authorizer

import (
	"errors"
	"regexp"
	"strings"

	"github.com/p1nant0m/gin-gin/authorization"
)

var (
	ErrRegoPackageSyntax = errors.New("rego package syntax is invalid")
)

type PolicyGetter interface {
	GetPolicy(key string) ([]string, error)
}

type Authorization struct {
	getter PolicyGetter
}

func NewAuthorization(getter PolicyGetter) authorization.AuthorizationInterface {
	return &Authorization{
		getter: getter,
	}
}

func (a *Authorization) GetModules(key string) (map[string]string, error) {
	modules, err := a.getter.GetPolicy(key)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`package\s.*`)
	regoMap := make(map[string]string)

	for _, regoCode := range modules {
		splitHeader := strings.Split(re.FindString(regoCode), " ")
		if len(splitHeader) < 1 {
			return nil, ErrRegoPackageSyntax
		}

		// Map: moduleName -> rego Rules definition
		regoMap[splitHeader[1]] = regoCode
	}

	return regoMap, nil
}
