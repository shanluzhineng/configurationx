package configurationx

type ConfigManager interface {
	Get(key string) ([]byte, error)
	List(key string) (KVPairs, error)
	Set(key string, value []byte) error
	Watch(key string, stop chan bool) <-chan *Response
}

type Response struct {
	Value []byte
	Error error
}

// 用来表示一个key/value值对
type KVPair struct {
	Key   string
	Value []byte
}

type KVPairs []*KVPair

// 一个用来从数据源中获取key/value值的接口
type Store interface {
	//获取指定key的值
	Get(Key string) ([]byte, error)
	//列取key中包含的所有子key列表
	List(key string) (KVPairs, error)
	//设置指定key的值
	Set(key string, value []byte) error
	//监控key的改变
	Watch(key string, stop chan bool) <-chan *Response
}
