package consulv

import (
	"fmt"
	"strings"
	"time"

	"github.com/abmpio/configurationx"
	"github.com/hashicorp/consul/api"
)

type ConsulClient struct {
	client    *api.KV
	waitIndex uint64
}

var _ configurationx.Store = (*ConsulClient)(nil)

// 构建一个consul实例
func NewConsulClient(machines []string) (*ConsulClient, error) {
	conf := api.DefaultConfig()
	if len(machines) > 0 {
		conf.Address = machines[0]
	}
	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return &ConsulClient{client.KV(), 0}, nil
}

// #region configurationx.Store Members

func (c *ConsulClient) Get(key string) ([]byte, error) {
	kv, _, err := c.client.Get(key, nil)
	if err != nil {
		return nil, err
	}
	if kv == nil {
		return nil, fmt.Errorf("key ( %s ) was not found", key)
	}
	return kv.Value, nil
}

func (c *ConsulClient) List(key string) (configurationx.KVPairs, error) {
	pairs, _, err := c.client.List(key, nil)
	if err != nil {
		return nil, err
	}
	ret := make(configurationx.KVPairs, len(pairs))
	for i, kv := range pairs {
		ret[i] = &configurationx.KVPair{Key: kv.Key, Value: kv.Value}
	}
	return ret, nil
}

func (c *ConsulClient) Set(key string, value []byte) error {
	key = strings.TrimPrefix(key, "/")
	kv := &api.KVPair{
		Key:   key,
		Value: value,
	}
	_, err := c.client.Put(kv, nil)
	return err
}

func (c *ConsulClient) Watch(key string, stop chan bool) <-chan *configurationx.Response {
	respChan := make(chan *configurationx.Response)
	go func() {
		for {
			opts := api.QueryOptions{
				WaitIndex: c.waitIndex,
			}
			keypair, meta, err := c.client.Get(key, &opts)
			if keypair == nil && err == nil {
				err = fmt.Errorf("key ( %s ) was not found", key)
			}
			if err != nil {
				respChan <- &configurationx.Response{
					Value: nil,
					Error: err,
				}
				time.Sleep(time.Second * 5)
				continue
			}
			c.waitIndex = meta.LastIndex
			respChan <- &configurationx.Response{
				Value: keypair.Value,
				Error: nil,
			}
		}
	}()
	return respChan
}

// #endregion
