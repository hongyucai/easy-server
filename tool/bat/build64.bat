set GOXMDIR="D:/phpstudy_pro/WWW/kaixinjishi/go-xm"
md %GOXMDIR%/build
cd %GOXMDIR%/build
go-bindata -pkg parse -o ../../../core/bindata/conf-data-live.go ../
rizla ../app/live/public/http.go
::go build  ../app/live/public/http.go
cd ..