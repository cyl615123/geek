package config

type Conf struct {
	Http  HttpConf
	Redis RedisConf
}

type HttpConf struct {
	Addr string
}

type RedisConf struct {
	Addr string
}
