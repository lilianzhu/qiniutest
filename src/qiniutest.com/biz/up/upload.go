package up

import (
	"qiniupkg.com/api.v7/kodo"
	. "qiniutest.com/configs"
)

func NewDefaultConfig() kodo.Config {
	cfg := kodo.Config{}

	cfg.AccessKey = ENV("access_key")
	cfg.SecretKey = ENV("secret_key")
	cfg.RSHost = ENV("rshost")
	cfg.RSFHost = ENV("rsfhost")
	cfg.IoHost = ENV("iohost")
	cfg.UpHosts = []string{ENV("uphost")}

	return cfg
}
