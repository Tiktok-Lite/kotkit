package etcd

import (
	"github.com/cloudwego/kitex/pkg/registry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	etcdRegister registry.Registry
)

func newEtcdRegistry() (registry.Registry, error) {
	_etcdRegister, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		logger.Errorf("Error occurs when creating etcd registry: %v", err)
		return nil, err
	}

	logger.Infof("Etcd registry start successfully at %s", etcdAddr)
	return _etcdRegister, nil
}

func Registry() (registry.Registry, error) {
	var err error
	once.Do(func() {
		etcdRegister, err = newEtcdRegistry()
	})

	if err != nil {
		return nil, err
	}

	return etcdRegister, nil
}
