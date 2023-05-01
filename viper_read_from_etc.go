package configurationx

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type GOOS string

const (
	GOOS_Android GOOS = "android"
	GOOS_Linux   GOOS = "linux"
	GOOS_Windows GOOS = "windows"

	ENV_AppName = "app.name"
)

var (
	_supportExtName map[string]bool = map[string]bool{
		"yaml": true,
		"yml":  true,
		"json": true,
	}
)

// 从etc目录的subPath下读取指定应用名称的配置信息
func (c *Configuration) readFromEtcFolder(subPath string, name string) *Configuration {
	if runtime.GOOS != string(GOOS_Linux) {
		// windows环境将不读取/etc目录中的配置"
		return c
	}
	if name == "" {
		name = getExecutableFileName()
		if len(name) <= 0 {
			return c
		}
	}
	if len(subPath) <= 0 {
		subPath = c.EtcSubPath
	}
	if len(subPath) <= 0 {
		return c
	}
	basePath := path.Join("/etc", c.EtcSubPath)
	for eachExtName := range _supportExtName {
		c.readAndMergeFile(basePath, name, eachExtName)
	}
	return c
}

func (c *Configuration) readFromWindowsUserFolder(subPath string, name string) *Configuration {
	if runtime.GOOS != string(GOOS_Windows) {
		return c
	}
	if name == "" {
		name = getExecutableFileName()
		if len(name) <= 0 {
			return c
		}
	}
	if len(subPath) <= 0 {
		subPath = c.EtcSubPath
	}
	if len(subPath) <= 0 {
		return c
	}
	userDir, err := os.UserHomeDir()
	if err != nil {
		return c
	}
	basePath := filepath.Join(userDir, c.EtcSubPath)
	for eachExtName := range _supportExtName {
		c.readAndMergeFile(basePath, name, eachExtName)
	}
	return c
}

func (c *Configuration) readAndMergeFile(basePath string, appName string, extName string) *Configuration {
	_, ok := _supportExtName[extName]
	if !ok {
		return c
	}
	fileFullpath := path.Join(basePath, appName+"."+extName)
	if _, err := os.Stat(fileFullpath); err != nil {
		// fmt.Println("默认配置文件不存在,不读取配置文件")
		return c
	}

	newViper := viper.New()
	newViper.SetConfigType(extName)
	newViper.AddConfigPath(basePath)
	newViper.SetConfigName(appName)
	err := newViper.ReadInConfig()
	if err != nil {
		err = fmt.Errorf("读取配置文件时出现异常,path:%s,异常信息:%s", fileFullpath, err.Error())
		panic(err)
	}
	c.viper.MergeConfigMap(newViper.AllSettings())
	return c
}

func getExecutableFileName() string {
	//先读取环境变量
	appName := os.Getenv(ENV_AppName)
	if len(appName) > 0 {
		return appName
	}
	path, _ := os.Executable()
	_, execFileName := filepath.Split(path)
	return execFileName
}

func ReadFromEtcFolder(fileName string) ConfigurationReadOption {
	return func(c *Configuration) {
		if runtime.GOOS != string(GOOS_Linux) {
			c.readFromWindowsUserFolder(c.EtcSubPath, fileName)
		} else {
			c.readFromEtcFolder(c.EtcSubPath, fileName)
		}
	}
}
