package loader
import (
	"github.com/kataras/golog"
	"go-xm/inits/bindata/conf"
	"gopkg.in/yaml.v2"
)

var (
	MongodbConfig Mongodb
)
func MongodbSettingParse() {
	//golog.Info("@@@ Init mongodb conf")
	mongodbData, err := parse.Asset("../config/mongodb.yml")
	if err != nil {
		golog.Fatalf("Error. %s", err)
	}
	if err = yaml.Unmarshal(mongodbData, &MongodbConfig); err != nil {
		golog.Fatalf("Error. %s", err)
	}
}
type Mongodb struct {
	Mongodbpro MongodbConfigInfo
	Mongodbdev MongodbConfigInfo
}

type MongodbConfigInfo struct {
	Username     string `yaml:"username"`
	Password string `yaml:"password"`
	Uri     string `yaml:"uri"`
	Database string `yaml:"database"`
	MaxOpenConns int `yaml:"maxOpenConns"`
	MaxIdleConns int `yaml:"maxIdleConns"`
}