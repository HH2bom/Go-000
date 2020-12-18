package appcontext

import (
	"app/internal/pkg/conf"
)

type Context struct {
	ops Options
}

func New(options ...Option) *Context {
	var ops Options
	for _, op := range options {
		op(ops)
	}

	return &Context{ops: ops}
}

func (c *Context) GetConf() conf.C {
	return c.ops.conf
}

type GetConf interface {
	GetConf() conf.C
}

type Options struct {
	conf conf.C
}
type Option func(ops Options)

func Config(c conf.C) Option {
	return func(ops Options) {
		ops.conf = c
	}
}
