package store

type PolicyStore interface {
	Lists() ([]string, error)
}
