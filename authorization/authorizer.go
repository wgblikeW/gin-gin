package authorization

import (
	"context"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	authv1 "github.com/p1nant0m/gin-gin/pkg/api/v1"
	"github.com/sirupsen/logrus"
)

const (
	KEY string = "authz-user"
)

type MarshalJSONData string

type Authorizer struct {
	engine *rego.PreparedEvalQuery
}

func NewAuthorizer(authorizationClient AuthorizationInterface) *Authorizer {
	var complieRego *ast.Compiler
	moduels, err := authorizationClient.GetModules(KEY)
	if err != nil {
		logrus.Fatalf("errors occurs when retrieve modules err:%+v\n", err)
	}

	complieRego, err = ast.CompileModules(moduels)
	if err != nil {
		logrus.Fatalf("errors occurs when compile modules err:%+v\n", err)
	}

	engine, err := rego.New(
		rego.Query("data.example.authz.allow"),
		rego.Compiler(complieRego),
	).PrepareForEval(context.TODO())
	if err != nil {
		logrus.Fatalf("errors occurs when create new rego engine err:%+v\n", err)
	}

	return &Authorizer{
		engine: &engine,
	}
}

func (a *Authorizer) Authorize(input interface{}) *authv1.Response {
	rs, err := a.engine.Eval(context.TODO(), rego.EvalInput(input))
	logrus.Debugf("evaluation: %+v", rs)

	var resp authv1.Response = authv1.Response{
		Denied: true,
		Reason: "input has been evaluated and was refused by the policy enforcement",
	}

	if err != nil {
		resp.Error = err.Error()
		resp.Reason = "errors occurs when evaluate the input"
		return &resp
	}

	if rs.Allowed() {
		resp.Allowed = true
		resp.Denied = false
		resp.Reason = "input is evaluated and authorized"
	}

	return &resp
}
