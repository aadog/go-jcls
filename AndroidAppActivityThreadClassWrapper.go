package cls

import (
	"github.com/aadog/go-java"
	"github.com/samber/mo"
)

type AndroidAppActivityThreadClassWrapper struct {
	*java.ClassWrapper
}

func (j *AndroidAppActivityThreadClassWrapper) CurrentActivityThread() mo.Result[*java.ObjectWrapper] {
	c, err := j.CallStaticObjectA("currentActivityThread").Get()
	if err != nil {
		return mo.Err[*java.ObjectWrapper](err)
	}
	return mo.Ok(java.ObjectWrapperWithJniPtr(c.JniPtr()))
}
func (j *AndroidAppActivityThreadClassWrapper) GetApplication() mo.Result[*java.ObjectWrapper] {
	c, err := j.CallStaticObjectA("getApplication").Get()
	if err != nil {
		return mo.Err[*java.ObjectWrapper](err)
	}
	return mo.Ok(java.ObjectWrapperWithJniPtr(c.JniPtr()))
}
func (j *AndroidAppActivityThreadClassWrapper) GetPackageManager() mo.Result[*java.ObjectWrapper] {
	c, err := j.CallStaticObjectA("getPackageManager").Get()
	if err != nil {
		return mo.Err[*java.ObjectWrapper](err)
	}
	return mo.Ok(java.ObjectWrapperWithJniPtr(c.JniPtr()))
}
func (j *AndroidAppActivityThreadClassWrapper) GetIntentBeingBroadcast() mo.Result[*java.ObjectWrapper] {
	c, err := j.CallStaticObjectA("getIntentBeingBroadcast").Get()
	if err != nil {
		return mo.Err[*java.ObjectWrapper](err)
	}
	return mo.Ok(java.ObjectWrapperWithJniPtr(c.JniPtr()))
}

func AndroidAppActivityThreadClass() *AndroidAppActivityThreadClassWrapper {
	return &AndroidAppActivityThreadClassWrapper{
		ClassWrapper: java.Use("android.app.ActivityThread").MustGet(),
	}
}
