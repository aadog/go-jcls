package cls

import (
	"github.com/aadog/go-java"
	"github.com/samber/mo"
)

type JavaIoBufferedReaderClassWrapper struct {
	*java.ClassWrapper
}

func (j *JavaIoBufferedReaderClassWrapper) BufferedReader_in(in *JavaIoInputStreamReaderObjectWrapper) mo.Result[*JavaIoBufferedReaderObjectWrapper] {
	c, err := j.New(in).Get()
	if err != nil {
		return mo.Err[*JavaIoBufferedReaderObjectWrapper](err)
	}
	return mo.Ok(JavaIoBufferedReaderWithJniPtr(c.JniPtr()))
}
func (j *JavaIoBufferedReaderClassWrapper) BufferedReader_in_sz(in *JavaIoInputStreamReaderObjectWrapper, sz int) mo.Result[*JavaIoBufferedReaderObjectWrapper] {
	c, err := j.New(in, sz).Get()
	if err != nil {
		return mo.Err[*JavaIoBufferedReaderObjectWrapper](err)
	}
	return mo.Ok(JavaIoBufferedReaderWithJniPtr(c.JniPtr()))
}

func JavaIoBufferedReaderClass() *JavaIoBufferedReaderClassWrapper {
	return &JavaIoBufferedReaderClassWrapper{
		ClassWrapper: java.Use("java.io.BufferedReader").MustGet(),
	}
}
