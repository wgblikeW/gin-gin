package authorization

type AuthorizationInterface interface {
	GetModules(string) (map[string]string, error)
}
