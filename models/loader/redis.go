package loader

import (
	"github.com/kataras/golog"
	"go-xm/inits/bindata/conf"
	"gopkg.in/yaml.v2"
)

var (
	RedisConfig Redis
)
func RedisSettingParse() {
	//golog.Info("@@@ Init redis conf")
	redisData, err := parse.Asset("../config/redis.yml")
	if err != nil {
		golog.Fatalf("Error. %s", err)
	}
	if err = yaml.Unmarshal(redisData, &RedisConfig); err != nil {
		golog.Fatalf("Error. %s", err)
	}
}
type Redis struct {
	Redispro RedisConfigInfo
	Redisdev RedisConfigInfo
}

type RedisConfigInfo struct {
	Addr  string `yaml:"addr"`
	Password string `yaml:"password"`
	Db     int `yaml:"db"`
}