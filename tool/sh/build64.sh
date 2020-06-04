cd /data/web/go-xm/build
go-bindata -pkg parse -o ../inits/bindata/conf/conf-data.go ../conf/
go build  ../cmd/http/http.go
go build  ../cmd/jd/GetJdOrdersDays.go
go build  ../cmd/jd/GetJdOrdersMinute.go
go build  ../cmd/jd/UpJdItems.go
go build  ../cmd/jd/UpJdCoupons.go
go build  ../cmd/jd/UpOrderassignFromMySelf.go
go build  ../cmd/jd/UpOrderassignFromJdordergoods.go
go build  ../cmd/cs/cs.go
cd ..