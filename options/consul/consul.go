package consul

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
)

const (
	defaultEnableHealthCheck              = true
	defaultHealthCheckInterval            = 10
	defaultHealthCheckTimeout             = 5
	defaultEnableHeartbeatCheck           = false
	defaultHeartbeatCheckInterval         = 10
	defaultDeregisterCriticalServiceAfter = 5
)

type ConsulOptions struct {
	Host       string `json:"host,omitempty"`
	Port       int    `json:"port,omitempty"`
	Datacenter string `json:"datacenter,omitempty"`
	AclToken   string `json:"aclToken,omitempty"`
	Disabled   bool   `json:"disabled,omitempty"`

	Registration *RegistrationInfo `json:"registration,omitempty"`
}

func (c *ConsulOptions) Normalize() *ConsulOptions {
	if c.Port <= 0 {
		c.Port = 8500
	}
	if c.Registration == nil {
		c.Registration = &RegistrationInfo{}
	}
	c.Registration.normalize()
	return c
}

// 注册到consul的参数
type RegistrationInfo struct {
	//是否启用注册，默认为false
	Enabled bool `json:"enabled,omitempty"`
	// 是否启用健康检查,默认为true
	EnableHealthCheck *bool `json:"enableHealthCheck,omitempty"`
	//健康检查时间间隔，默认10秒
	HealthCheckInterval *int `json:"healthCheckInterval,omitempty"`
	//健康检查超时时间，默认5秒
	HealthCheckTimeout *int   `json:"healthCheckTimeout,omitempty"`
	HealthCheckHTTP    string `json:"healthCheckHTTP,omitempty"`
	//告诉注册中心使用TCP来进行服务健康检查,格式为 （IP地址:port）类型的值，如127.0.0.1:8000
	HealthCheckTCP string `json:"healthCheckTCP,omitempty"`
	//是否启用心跳检查,默认为true
	EnableHeartbeatCheck *bool `json:"enableHeartbeatCheck,omitempty"`
	//心跳检查时间间隔，默认10秒
	HeartbeatCheckInterval *int `json:"heartbeatCheckInterval,omitempty"`
	//健康检测失败后自动注销服务时间，默认5秒
	DeregisterCriticalServiceAfter *int `json:"deregisterCriticalServiceAfter,omitempty"`

	ID string `json:"id,omitempty"`
	//注册到注册中心的服务所属的产品，默认值为程序名
	Product string `json:"product,omitempty"`
	//服务名
	ServiceName string `json:"serviceName,omitempty"`
	//注册到注册中心的服务地址列表,一个有效的url格式
	//tcp协议如下:tcp://host:port
	//http/https协议如下:http://host:port/path
	Endpoint []string `json:"endpoint,omitempty"`
	//tags
	Tags []string `json:"tags,omitempty"`
	//注册到注册中心时给的Meta属性值
	Meta map[string]string `json:"meta,omitempty"`
}

type ServiceAddressUrl struct {
	Scheme  string
	Address string
	Port    int
}

// 将Endpoint属性解析成一组服务，key为url的scheme,如http,tcp等
func (r *RegistrationInfo) ParseServiceAddress() (map[string]*ServiceAddressUrl, error) {
	if len(r.Endpoint) <= 0 {
		return nil, nil
	}
	serviceAddressList := make(map[string]*ServiceAddressUrl)
	for _, eachEndpoint := range r.Endpoint {
		urlValue, err := url.Parse(eachEndpoint)
		if err != nil {
			return nil, fmt.Errorf("无效的endpoint值,必须为一个有效的url")
		}
		host, portValue, err := net.SplitHostPort(urlValue.Host)
		if err != nil {
			return nil, err
		}
		port, err := strconv.Atoi(portValue)
		if err != nil {
			return nil, err
		}
		serviceAddressList[urlValue.Scheme] = &ServiceAddressUrl{
			Scheme:  urlValue.Scheme,
			Address: host,
			Port:    port,
		}
	}
	return serviceAddressList, nil
}

func (r *RegistrationInfo) ParseServiceAddressForScheme(scheme string) (*ServiceAddressUrl, error) {
	list, err := r.ParseServiceAddress()
	if err != nil {
		return nil, err
	}
	serviceAddress, ok := list[scheme]
	if !ok {
		return nil, nil
	}
	return serviceAddress, nil
}

// 检测是否需要做通
func (r *RegistrationInfo) normalize() *RegistrationInfo {
	if r.Meta == nil {
		r.Meta = make(map[string]string)
	}
	r.EnableHealthCheck = makeBoolPtr(r.EnableHealthCheck, defaultEnableHealthCheck)
	r.HealthCheckInterval = makeIntPtr(r.HealthCheckInterval, defaultHealthCheckInterval)
	r.HealthCheckTimeout = makeIntPtr(r.HealthCheckTimeout, defaultHealthCheckTimeout)
	r.EnableHeartbeatCheck = makeBoolPtr(r.EnableHeartbeatCheck, defaultEnableHeartbeatCheck)
	r.HeartbeatCheckInterval = makeIntPtr(r.HealthCheckTimeout, defaultHeartbeatCheckInterval)
	r.DeregisterCriticalServiceAfter = makeIntPtr(r.DeregisterCriticalServiceAfter, defaultDeregisterCriticalServiceAfter)
	return r
}
