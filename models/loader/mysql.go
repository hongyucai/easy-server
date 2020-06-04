package loader

import (
	"github.com/kataras/golog"
	"go-xm/inits/bindata/conf"
	"gopkg.in/yaml.v2"
)

var (
	MysqlConfig Mysql
)
func MysqlSettingParse() {
	//golog.Info("@@@ Init mysql conf")
	mysqlData, err := parse.Asset("../config/mysql.yml")
	if err != nil {
		golog.Fatalf("Error. %s", err)
	}
	if err = yaml.Unmarshal(mysqlData, &MysqlConfig); err != nil {
		golog.Fatalf("Error. %s", err)
	}
}
type Mysql struct {
	Mysqlpro MysqlConfigInfo
	Mysqldev MysqlConfigInfo
}

type MysqlConfigInfo struct {
	Dialect  string `yaml:"dialect"`
	Username     string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
	ShowSql  bool   `yaml:"showSql"`
	LogLevel string `yaml:"logLevel"`
	MaxOpenConns int `yaml:"maxOpenConns"`
	MaxIdleConns int `yaml:"maxIdleConns"`
}