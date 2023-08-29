package etcd

import (
	"github.com/cloudwego/kitex/pkg/discovery"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	etcdResolver discovery.Resolver
)

func newResolver() (discovery.Resolver, error) {
	_etcdResolver, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		logger.Errorf("Error occurs when creating etcd resolver: %v", err)
		return nil, err
	}

	logger.Infof("Etcd resolver start successfully at %s", etcdAddr)
	return _etcdResolver, nil
}

func Resolver() (discovery.Resolver, error) {
	var err error
	once.Do(func() {
		etcdResolver, err = newResolver()
	})

	if err != nil {
		return nil, err
	}

	return etcdResolver, nil
}
