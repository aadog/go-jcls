package cls

import (
	"github.com/aadog/go-java"
	"github.com/samber/mo"
)

type JavaIoInputStreamReaderClassWrapper struct {
	*java.ClassWrapper
}

func (j *JavaIoInputStreamReaderClassWrapper) InputStreamReader_in(in *JavaIoInputStreamObjectWrapper) mo.Result[*JavaIoInputStreamReaderObjectWrapper] {
	c, err := j.New(in).Get()
	if err != nil {
		return mo.Err[*JavaIoInputStreamReaderObjectWrapper](err)
	}
	return mo.Ok(JavaIoInputStreamReaderWithJniPtr(c.JniPtr()))
}
func (j *JavaIoInputStreamReaderClassWrapper) InputStreamReader_in_cs(in *JavaIoInputStreamObjectWrapper, cs *jvm.ObjectWrapper) mo.Result[*JavaIoInputStreamReaderObjectWrapper] {
	c, err := j.New(in, cs).Get()
	if err != nil {
		return mo.Err[*JavaIoInputStreamReaderObjectWrapper](err)
	}
	return mo.Ok(JavaIoInputStreamReaderWithJniPtr(c.JniPtr()))
}
func (j *JavaIoInputStreamReaderClassWrapper) InputStreamReader_in_dec(in *JavaIoInputStreamObjectWrapper, dec *jvm.ObjectWrapper) mo.Result[*JavaIoInputStreamReaderObjectWrapper] {
	c, err := j.New(in, dec).Get()
	if err != nil {
		return mo.Err[*JavaIoInputStreamReaderObjectWrapper](err)
	}
	return mo.Ok(JavaIoInputStreamReaderWithJniPtr(c.JniPtr()))
}
func (j *JavaIoInputStreamReaderClassWrapper) InputStreamReader_in_charsetName(in *JavaIoInputStreamObjectWrapper, charsetName string) mo.Result[*JavaIoInputStreamReaderObjectWrapper] {
	c, err := j.New(in, charsetName).Get()
	if err != nil {
		return mo.Err[*JavaIoInputStreamReaderObjectWrapper](err)
	}
	return mo.Ok(JavaIoInputStreamReaderWithJniPtr(c.JniPtr()))
}

func JavaIoInputStreamReaderClass() *JavaIoInputStreamReaderClassWrapper {
	return &JavaIoInputStreamReaderClassWrapper{
		ClassWrapper: java.Use("java.io.InputStreamReader").MustGet(),
	}
}
