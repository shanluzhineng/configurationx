package consulv

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/abmpio/configurationx"
	"github.com/abmpio/configurationx/options/consul"
)

func ReadFromConsul(consulOptions consul.ConsulOptions, consulPathList []string) *configurationx.Configuration {
	c := configurationx.New()
	if len(consulPathList) <= 0 {
		return c
	}
	c.BaseConsulPathList = consulPathList
	if len(consulOptions.Host) <= 0 {
		c.Logger.Info("没有配置好consul.host参数,不从consul中读取配置")
		return c
	}

	if consulOptions.Disabled {
		c.Logger.Info("consul.disabled参数值为true,将不从consul中读取配置")
		return c
	}
	endpoint := fmt.Sprintf("%s:%d", consulOptions.Host, consulOptions.Port)
	err := initConfigManager(c, []string{endpoint})
	if err != nil {
		return c
	}

	childKeyList, err := getChildKvPairs(c, c.BaseConsulPathList, true)
	if err != nil {
		panic(err)
	}
	if len(childKeyList) <= 0 {
		//没有子节点，则直接返回
		return c
	}

	for _, eachChildKey := range childKeyList {
		allSettings, err := getDataFromConfigManager(c, c.ConfigManager, eachChildKey)
		if err != nil {
			continue
		}
		//合并配置
		c.GetViper().MergeConfigMap(allSettings)
	}
	return c
}

func getChildKvPairs(c *configurationx.Configuration, keyList []string, containSubPath bool) (childKeyList []string, err error) {
	if len(keyList) <= 0 {
		return childKeyList, nil
	}
	for _, eachKey := range keyList {
		thisChildKvPairs, err := c.ConfigManager.List(eachKey)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("获取%s的子key时出现异常,详细异常信息:%s", eachKey, err.Error()))
			return childKeyList, err
		}
		if len(thisChildKvPairs) <= 0 {
			continue
		}
		for _, eachChildKey := range thisChildKvPairs {
			if eachChildKey == nil {
				continue
			}
			if eachChildKey.Key == eachKey+"/" {
				//自身，直接循环
				continue
			}
			if strings.HasSuffix(eachChildKey.Key, "/") && containSubPath {
				//递归获取子节点
				subChildKeyList, err := getChildKvPairs(c, []string{eachChildKey.Key}, containSubPath)
				if err != nil {
					return childKeyList, err
				}
				if len(subChildKeyList) <= 0 {
					continue
				}
				childKeyList = append(childKeyList, subChildKeyList...)
			} else {
				childKeyList = append(childKeyList, eachChildKey.Key)
			}
		}
	}
	return childKeyList, nil
}

func initConfigManager(c *configurationx.Configuration, endPoint []string) error {
	configManager, err := NewStandardConsulConfigManager(endPoint)
	if err != nil {
		c.Logger.Error("连接到consul时出现异常,异常信息:%s", err.Error())
		return err
	}
	c.ConfigManager = configManager
	return nil
}

func getDataFromConfigManager(c *configurationx.Configuration, cm configurationx.ConfigManager, path string) (map[string]interface{}, error) {
	b, err := cm.Get(path)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(b, &result)
	if err != nil {
		c.Logger.Error("cann't unmarshal map from data, path: %s", path)
		return nil, err
	}
	return result, err
}
