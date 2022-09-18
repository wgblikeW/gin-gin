package store

type Factory interface {
	Policy() PolicyStore
}
