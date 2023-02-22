package consul

import (
	"strconv"
	"time"
)

const (
	//用来描述实例的产品的meta key
	MetaName_Product = "product"
	//当注册到consul时，所有基于abmp的框架的都会在meta中加入此属性，以指定注册到consul中的这个服务是中台的实例
	MetaName_IsHostInABMP = "isHostInABMP"
	//用来描述实例的描述信息的meta key
	MetaName_Description = "description"
	//用来描述实例的启动时间的meta key
	MetaName_StartTime = "startTime"
	//用来描述实例的应用版本号的meta key
	MetaName_AppVersion = "appVersion"
	//用来描述实例的中台框架版本号的meta key
	MetaName_AppFrameworkVersion = "appFrameworkVersion"
	//用来描述实例的运行环境的meta key,值主要有windows,supervisor,docker,systemd,other
	MetaName_HostEnvironment = "hostEnvironment"
	//用来描述实例所提供的web api的http地址的meta key
	MetaName_Http = "http"
	//用来描述实例的站点地址的meta key
	MetaName_Website = "website"
	//用来描述实例所提供的web api的用于consul健康检查的地址
	MetaName_Healthcheck = "healthcheck"
)

func (r *RegistrationInfo) SetMeta_Product(product string) *RegistrationInfo {
	return r.setMetaValue(MetaName_Product, product)
}

func (r *RegistrationInfo) SetMeta_IsHostInABMP(value bool) *RegistrationInfo {
	return r.setMetaValue(MetaName_IsHostInABMP, strconv.FormatBool(value))
}

func (r *RegistrationInfo) SetMeta_Description(value string) *RegistrationInfo {
	return r.setMetaValue(MetaName_Description, value)
}

func (r *RegistrationInfo) SetMeta_StartTime(t time.Time) *RegistrationInfo {
	return r.setMetaValue(MetaName_StartTime, t.Format(time.RFC3339))
}

func (r *RegistrationInfo) SetMeta_AppVersion(v string) *RegistrationInfo {
	return r.setMetaValue(MetaName_AppVersion, v)
}

func (r *RegistrationInfo) SetMeta_AppFrameworkVersion(v string) *RegistrationInfo {
	return r.setMetaValue(MetaName_AppFrameworkVersion, v)
}

func (r *RegistrationInfo) SetMeta_HostEnvironment(v string) *RegistrationInfo {
	return r.setMetaValue(MetaName_HostEnvironment, v)
}

func (r *RegistrationInfo) SetMeta_Http(v string) *RegistrationInfo {
	return r.setMetaValue(MetaName_Http, v)
}

func (r *RegistrationInfo) SetMeta_Website(v string) *RegistrationInfo {
	return r.setMetaValue(MetaName_Website, v)
}

func (r *RegistrationInfo) SetMeta_Healthcheck(v string) *RegistrationInfo {
	return r.setMetaValue(MetaName_Healthcheck, v)
}

func (r *RegistrationInfo) setMetaValue(key string, v string) *RegistrationInfo {
	r.Meta[key] = v
	return r
}
