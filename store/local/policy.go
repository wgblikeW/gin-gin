package local

import "errors"

type policy struct {
	op *LocalStorage
}

var (
	ErrPolicyNotFound = errors.New("policy not found")
)

func newPolicies(op *LocalStorage) *policy {
	return &policy{op}
}

func (p *policy) Lists() ([]string, error) {
	p.op.lock.RLock()
	defer p.op.lock.RUnlock()

	if item, exists := p.op.policies["authz-user"]; exists {
		return item, nil
	}

	return nil, ErrPolicyNotFound
}
