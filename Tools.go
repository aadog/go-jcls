package jcls

import (
	"github.com/aadog/go-java"
	"github.com/samber/mo"
)

var Tools ToolsStruct

type ToolsStruct struct {
}

// GetApplicationContext request androidContext
func (t ToolsStruct) GetApplicationContext() mo.Result[*java.ObjectWrapper] {
	androidAppActivityClass := AndroidAppActivityThreadClass()
	currentActivityThread := androidAppActivityClass.CallStaticObjectA("currentActivityThread").MustGet()
	defer currentActivityThread.Free()
	application := currentActivityThread.CallObjectA("getApplication").MustGet()
	defer application.Free()
	applicationContext := application.CallObjectA("getApplicationContext").MustGet()
	return mo.Ok(applicationContext)
}
func (t ToolsStruct) GetPackageName(applicationContext *java.ObjectWrapper) mo.Result[string] {
	packageName := applicationContext.CallStringA("getPackageName").MustGet()
	return mo.Ok(packageName)
}
func (t ToolsStruct) GetVersionName(applicationContext *java.ObjectWrapper) mo.Result[string] {
	packageName := t.GetPackageName(applicationContext).MustGet()
	pkgManager := applicationContext.CallObjectA("getPackageManager").MustGet()
	defer pkgManager.Free()
	pkgInfo := pkgManager.CallObjectA("getPackageInfo", packageName, 0).MustGet()
	defer pkgInfo.Free()
	str := pkgInfo.GetStringFieldValue("versionName").MustGet()
	return mo.Ok(str)
}

func (t ToolsStruct) GetNativeLibraryDir(application *java.ObjectWrapper) mo.Result[string] {
	appInfo, err := application.CallObjectA("getApplicationInfo").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	defer appInfo.Free()
	return appInfo.GetStringFieldValue("nativeLibraryDir")
}
func (t ToolsStruct) GetGetProperty(property string) mo.Result[string] {
	system, err := java.Use("java.lang.System").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return system.CallStaticStringA("getProperty", property)
}
