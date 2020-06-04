package libs

import (
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"go-xm/inits"
	"go-xm/inits/parse"
	"golang.org/x/net/context"
)
//https://developer.qiniu.com/kodo/sdk/1238/go#returnbody-uptoken
func qiniuConf() parse.QiniuConfigInfo {
	if inits.QiniuPro {
		return parse.QiniuConfig.Qiniupro
	}else {
		return parse.QiniuConfig.Qiniudev
	}
}
func getUpToken() string {
	qiniu :=qiniuConf()
	putPolicy := storage.PutPolicy{
		Scope: qiniu.Bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(qiniu.AccessKey,qiniu.SecretKey)
	return putPolicy.UploadToken(mac)
}
//数据流上传（表单方式）
func upIoFile (key,localFile string) string {
	qiniu :=qiniuConf()
	putPolicy := storage.PutPolicy{
		Scope: qiniu.Bucket,
	}
	mac := qbox.NewMac(qiniu.AccessKey,qiniu.SecretKey)
	upToken :=  putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	resumeUploader := storage.NewResumeUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputExtra{
	}
	err := resumeUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//fmt.Println(ret.Key, ret.Hash)
	return qiniu.HttpsUrl+"/"+ret.Key
}
