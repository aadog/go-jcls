package jcls

import (
	"github.com/aadog/go-java"
	"github.com/samber/mo"
)

type JavaNetUrlClassWrapper struct {
	*java.ClassWrapper
}

func (j *JavaNetUrlClassWrapper) URL_spec(spec string) mo.Result[*JavaNetUrlObjectWrapper] {
	c, err := j.New(spec).Get()
	if err != nil {
		return mo.Err[*JavaNetUrlObjectWrapper](err)
	}
	return mo.Ok(JavaNetUrlObjectWrapperWithJniPtr(c.JniPtr()))
}
func (j *JavaNetUrlClassWrapper) URL_protocol_host_port_file(protocol string, host string, port int, file string) mo.Result[*JavaNetUrlObjectWrapper] {
	c, err := j.New(protocol, host, port, file).Get()
	if err != nil {
		return mo.Err[*JavaNetUrlObjectWrapper](err)
	}
	return mo.Ok(JavaNetUrlObjectWrapperWithJniPtr(c.JniPtr()))
}

func JavaNetUrlClass() *JavaNetUrlClassWrapper {
	return &JavaNetUrlClassWrapper{
		ClassWrapper: java.Use("java.net.URL").MustGet(),
	}
}
