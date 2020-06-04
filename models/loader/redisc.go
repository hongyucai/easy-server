package loader
import (
	"github.com/kataras/golog"
	"go-xm/inits/bindata/conf"
	"gopkg.in/yaml.v2"
)

var (
	RediscConfig Redisc
)
func RediscSettingParse() {
	//golog.Info("@@@ Init redisc conf")
	rediscData, err := parse.Asset("../config/redisc.yml")
	if err != nil {
		golog.Fatalf("Error. %s", err)
	}
	if err = yaml.Unmarshal(rediscData, &RediscConfig); err != nil {
		golog.Fatalf("Error. %s", err)
	}
}
type Redisc struct {
	Rediscpro RediscConfigInfo
	Rediscdev RediscConfigInfo
}

type RediscConfigInfo struct {
	Addr1  string `yaml:"addr1"`
	Addr2  string `yaml:"addr2"`
	Addr3  string `yaml:"addr3"`
	Addr4  string `yaml:"addr4"`
	Addr5  string `yaml:"addr5"`
	Addr6  string `yaml:"addr6"`
}