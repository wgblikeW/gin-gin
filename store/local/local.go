package local

import (
	"errors"
	"sync"

	_ "embed"

	"github.com/p1nant0m/gin-gin/store"
	"github.com/sirupsen/logrus"
)

type datastore struct {
	ls *LocalStorage
}

func (ds *datastore) Policy() store.PolicyStore {
	return newPolicies(ds.ls)
}

var (
	//go:embed rules/policy1.rego
	policy1 string

	localStorageFactory store.Factory
)

type LocalStorage struct {
	lock     *sync.RWMutex
	policies map[string][]string
}

var (
	onceCreate sync.Once
)

func GetLocalFactoryOrExit() store.Factory {
	onceCreate.Do(func() {
		var policies map[string][]string = make(map[string][]string)
		policies["authz-user"] = append(policies["authz-user"], policy1)

		ls := &LocalStorage{
			lock:     new(sync.RWMutex),
			policies: policies,
		}

		localStorageFactory = &datastore{ls}
		logrus.Info("localStorage has been created successfully")
	})

	if localStorageFactory == nil {
		logrus.Fatal("failed to get localStorage store factory")
	}

	return localStorageFactory
}

type DirectLocalStorage struct {
	ls store.Factory
}

func GetDirectLocalInsOr(store store.Factory) (*DirectLocalStorage, error) {
	if store != nil {
		return &DirectLocalStorage{store}, nil
	}

	return nil, errors.New("unexpected errors GetLocalInsOr receive nil store.Factory")
}

func (dls *DirectLocalStorage) GetPolicy(key string) ([]string, error) {
	return dls.ls.Policy().Lists()
}
