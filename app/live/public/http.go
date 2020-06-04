package main

import (
	"flag"
	"fmt"
	"github.com/kataras/iris/v12"
	conf "go-mod/app/live/config"
	"go-mod/core/providers"
	"net/http"
	_ "net/http/pprof"
	"log"
)

func init()  {
	fmt.Printf("go init");
}

//# 安装rizla包
//$ go get -u github.com/kataras/rizla
//# 热重启方式启动iris项目
//$ rizla http.go
//rizla http.go //#单个项目监视
//rizla C:/myprojects/project1/main.go C:/myprojects/project2/main.go //＃多项目监控
//rizla -walk main.go //#仅在默认文件更改时才在扫描之前添加“ -walk”选项对您不起作用。
//rizla -delay=5s main.go //#如果delay> 0，那么它将延迟重新加载，还请注意，它接受第一个更改，但其余的每个更改都“延迟”。

// $ go run http.go  go build http.go
const maxSize = 5 << 20 //限制上传5MB
func main() {
	//性能监测
	flag.Parse()
	go func() {
		log.Println(http.ListenAndServe("localhost:6061", nil))
	}()
	//性能检测结束代码
	app := iris.New()
	providers.Hub(app)
	//app.Run(iris.TLS("127.0.0.1:443", "mycert.cert", "mykey.key"))
	app.Run(iris.Addr(":" + conf.O.Port), iris.WithConfiguration(conf.C),iris.WithPostMaxMemory(maxSize))
}