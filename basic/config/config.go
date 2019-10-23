package config

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	"sync"
)

var (
	m      sync.RWMutex
	inited bool
	c      = &configurator{}
)

type Configurator interface {
	App(name string, config interface{}) (err error)
}

type configurator struct {
	conf config.Config
}

func (c *configurator) App(name string, config interface{}) (err error) {
	v := c.conf.Get(name)
	if v != nil {
		err = v.Scan(config)
	} else {
		err = fmt.Errorf("[App] conf not exists ,%s", name)
	}

	return
}

func C() Configurator {
	return c
}

func (c *configurator) init(ops Options) (err error) {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Logf("[init] config already inited ")
		return
	}

	c.conf = config.NewConfig()
	err = c.conf.Load(ops.Sources...)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Logf("[init] watch config change ")

		watcher, err := c.conf.Watch()
		if err != nil {
			log.Fatal(err)
		}

		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatal(err)
			}

			log.Logf("[init] conf changed ,%v", string(v.Bytes()))
		}
	}()

	inited = true
	return
}

func Init(opts ...Option) {
	ops := Options{}
	for _, o := range opts {
		o(&ops)
	}

	c = &configurator{}
	c.init(ops)
}
