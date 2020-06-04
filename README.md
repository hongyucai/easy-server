
│  .gitignore
│  go.mod
│  go.sum
│  README.md
│ 
├─app                                   项目层
│  └─live
│      │  .env
│      │  
│      ├─config
│      ├─http
│      │  ├─api
│      │  │      room.go
│      │  │      
│      │  ├─controllers
│      │  │      LiveController.go
│      │  │      
│      │  └─mapper
│      ├─resources
│      ├─routes
│      │      http.go
│      │      
│      └─storage
├─auth
│  └─jwts
│          jwts.go
│          middleware.go
│          
├─common
│  ├─helper
│  │      aes.go
│  │      encrypt.go
│  │      http.go
│  │      http_test.go
│  │      snowFlake.go
│  │      sysconf.go
│  │      times.go
│  │      util.go
│  │      
│  └─libs
│          jd.go
│          pdd.go
│          qiniu.go
│          robot.go
│          wechat.go
│          
├─gateway
├─job
├─models                                数据层
│  ├─config
│  │      mongodb.yml
│  │      mysql.yml
│  │      redis.yml
│  │      redisc.yml
│  │      
│  ├─loader
│  │      mongodb.go
│  │      mysql.go
│  │      redis.go
│  │      redisc.go
│  │      
│  ├─migrations
│  └─sources
│      ├─mongodb
│      ├─mysql
│      └─redis
├─resources
├─server                        服务层 商品服务、订单服务
├─test
└─tool
    ├─bat
    │      build64.bat
    │      
    └─sh
            build64.sh
            guards.sh
            
			
go run p2g.go D:\phpstudy_pro\WWW\kaixinjishi\we7\addons\shimmer_liveshop\app\Modules\Liveapi\Http\Controllers\shop\GoodsController.php
