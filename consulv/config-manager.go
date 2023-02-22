package consulv

import (
	"github.com/abmpio/configurationx"
)

type standardConfigManager struct {
	store configurationx.Store
}

// 创建一个ConfigManager的新实例
func NewStandardConfigManager(client configurationx.Store) (configurationx.ConfigManager, error) {
	return standardConfigManager{client}, nil
}

var _ configurationx.ConfigManager = (*standardConfigManager)(nil)

// #region configurationx.ConfigManager Members

func (c standardConfigManager) Get(key string) ([]byte, error) {
	return c.store.Get(key)
}

func (c standardConfigManager) List(key string) (configurationx.KVPairs, error) {
	return c.store.List(key)
}

func (c standardConfigManager) Set(key string, value []byte) error {
	return c.store.Set(key, value)
}

func (c standardConfigManager) Watch(key string, stop chan bool) <-chan *configurationx.Response {
	resp := make(chan *configurationx.Response)
	backendResp := c.store.Watch(key, stop)
	go func() {
		for {
			select {
			case <-stop:
				return
			case r := <-backendResp:
				if r.Error != nil {
					resp <- &configurationx.Response{
						Value: nil,
						Error: r.Error,
					}
					continue
				}
				resp <- &configurationx.Response{
					Value: r.Value,
					Error: nil,
				}
			}
		}
	}()
	return resp
}

// #endregion

// 返回一个通过consul来读取的配置对象
func NewStandardConsulConfigManager(machines []string) (configurationx.ConfigManager, error) {
	store, err := NewConsulClient(machines)
	if err != nil {
		return nil, err
	}
	return NewStandardConfigManager(store)
}
