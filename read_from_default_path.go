package configurationx

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

const (
	defaultConfigFile string = "config"
)

var (
	supportedFileExtList = []string{".yaml", ".yml", ".json"}
)

// // 读取配置文件
// func (c *Configuration) ReadConfiguration(path ...string) *Configuration {
// 	// fmt.Println("准备读取应用程序目录下的./etc/config.yml配置文件")
// 	consulViper := viper.New()
// 	setupViperFromDefaultPath(consulViper)
// 	c.ReadFromEtcFolder()
// 	consulViper.MergeConfigMap(c.viper.AllSettings())
// 	consul.ReadFrom(consulViper)

// 	return c
// }

func setupViperFromDefaultPath(v *viper.Viper) {
	basePath, err := os.Getwd()
	if err != nil {
		err = fmt.Errorf("os.GetWd error, err:%v", err)
		panic(err)
	}
	filePath := path.Join(basePath, "etc")
	fileList, _ := discoverFileFromPath(filePath, supportedFileExtList)
	if len(fileList) <= 0 {
		// empty folder
		return
	}
	for _, eachFile := range fileList {
		fileName := path.Base(eachFile)
		i := strings.LastIndex(fileName, ".")
		if i == -1 {
			continue
		}
		configName := fileName[:i]
		configType := fileName[i+1:]
		if configName == "" || configType == "" {
			continue
		}
		v.SetConfigType(configType)
		v.SetConfigName(configName)
		v.AddConfigPath(filePath)
		err := v.ReadInConfig()
		if err != nil {
			err = fmt.Errorf("读取配置文件时出现异常,文件名:%s,异常信息:%s", eachFile, err.Error())
			panic(err)
		}
	}
}

func (c *Configuration) readFromDefaultPath() {
	defaultViper := viper.New()
	setupViperFromDefaultPath(defaultViper)
	//合并
	c.viper.MergeConfigMap(defaultViper.AllSettings())
}

// ConfigurationReadOption that read from ./etc path, belown file type is searched,
// 1. *.yml
// 2. *.yaml
// 3. *.json
func ReadFromDefaultPath() ConfigurationReadOption {
	return func(c *Configuration) {
		c.readFromDefaultPath()
	}
}
