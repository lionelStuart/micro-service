package basic

import "micro-service/basic/config"

var (
	pluginFuncs []func()
)

type Options struct {
	EnabledDB    bool
	EnabledRedis bool
	cfgOps       []config.Option
}

type Option func(o *Options)

func Init(opts ...config.Option) {
	config.Init(opts...)

	for _, f := range pluginFuncs {
		f()
	}
}

func Register(f func()) {
	pluginFuncs = append(pluginFuncs, f)
}
