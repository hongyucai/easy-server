package libs

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/golog"
	"go-xm/inits"
	"go-xm/inits/parse"
	"go-xm/utils"
	"strings"
)
func wechatxcxConf() parse.WechatXcxConfigInfo {
	if inits.WechatPro {
		return parse.WechatConfig.Wechatxcxpro
	}else {
		return parse.WechatConfig.Wechatxcxdev
	}
}
func wechatgzhConf() parse.WechatGzhConfigInfo {
	if inits.WechatPro {
		return parse.WechatConfig.Wechatgzhpro
	}else {
		return parse.WechatConfig.Wechatgzhdev
	}
}
func getAccessTokenGzh() string { //[]byte
	var (
		url ="https://api.weixin.qq.com/cgi-bin/token"
		query = map[string]string{}
	)
	query = map[string]string{
		"appid":wechatgzhConf().Appid,
		"secret":wechatgzhConf().Secret,
		"grant_type":"client_credential",
	}
	//GET
	h := utils.NewHttpSend(utils.GetUrlBuild(url,query))
	strb, err := h.Get()
	if err != nil {
		golog.Errorf("wx: getAccessTokenGzh http get err: %s", err.Error())
		return ""
	}
	//strs := strings.Replace(string(strb),"\\", "",-1)
	fmt.Printf(string(strb))
	// []byte(strs) //string 转  byte
	// string(strb[:]) //byte 转  string
	resultJson, err1 := simplejson.NewJson(strb)
	if err1 != nil {
		golog.Errorf("@@@ sessionKey %s", err1)
		return ""
	}
	//fmt.Printf(string(jsonStr))
	return resultJson.Get("access_token").MustString()
}
func getAccessTokenXcx() string { //[]byte
	var (
		url ="https://api.weixin.qq.com/cgi-bin/token"
		query = map[string]string{}
	)
	query = map[string]string{
		"appid":wechatxcxConf().Appid,
		"secret":wechatxcxConf().Secret,
		"grant_type":"client_credential",
	}
	//GET
	h := utils.NewHttpSend(utils.GetUrlBuild(url,query))
	strb, err := h.Get()
	if err != nil {
		golog.Errorf("wx: getAccessToken http get err: %s", err.Error())
		return ""
	}
	//strs := strings.Replace(string(strb),"\\", "",-1)
	fmt.Printf(string(strb))
	// []byte(strs) //string 转  byte
	// string(strb[:]) //byte 转  string
	resultJson, err1 := simplejson.NewJson(strb)
	if err1 != nil {
		golog.Errorf("@@@ sessionKey %s", err1)
		return ""
	}
	//fmt.Printf(string(jsonStr))
	return resultJson.Get("access_token").MustString()
}
func sessionKey(code string) map[string]interface{} {
	var (
		url ="https://api.weixin.qq.com/sns/jscode2session"
		query = map[string]string{}
		data = map[string]interface{}{}
	)
	query = map[string]string{
		"appid":wechatxcxConf().Appid,
		"secret":wechatxcxConf().Secret,
		"grant_type":"authorization_code",
		"js_code":code,
	}
	h := utils.NewHttpSend(utils.GetUrlBuild(url,query))
	strb, err := h.Get()
	if err != nil {
		golog.Errorf("wx: sessionKey http get err: %s", err.Error())
		return nil
	}
	resultJson, err1 := simplejson.NewJson(strb)
	if err1 != nil {
		golog.Errorf("@@@ sessionKey %s", err1)
		return nil
	}
	//fmt.Printf(string(jsonStr))
	data["openid"] = resultJson.Get("openid").MustString()
	data["session_key"] = resultJson.Get("session_key").MustString()
	data["unionid"] = resultJson.Get("unionid").MustString()
	data["errcode"] = resultJson.Get("errcode").MustInt()
	data["errmsg"] = resultJson.Get("errmsg").MustMap()
	return data
}

func getQrcode(appName,scene,page string,width int64,name string) string {
	var (
		url ="https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token="+getAccessTokenXcx()
		json = map[string]interface{}{}
	)
	json["scene"]=scene
	json["width"]=string(width)
	json["page"]=page

	h := utils.NewHttpSend(url)
	h.SetSendType("JSON")
	h.SetBody(json)
	strb, err := h.Post()
	if err != nil {
		golog.Errorf("wx: getQrcode http Post err: %s", err.Error())
		return ""
	}
	if strings.Index(string(strb[:]), "errcode") != -1 {
		return ""
	}
	return  upIoFile(name,string(strb[:]))
}
func uploadWxMedia(appName,name,path string) string {
	var (
		url ="https://api.weixin.qq.com/cgi-bin/media/upload?type=image&access_token="+getAccessTokenXcx()
	)
	strb, err := utils.PostFile(url,path,"media")
	if err != nil {
		golog.Errorf("wx: uploadWxMedia http Post err: %s", err.Error())
		return ""
	}
	resultJson, err1 := simplejson.NewJson(strb)
	if err1 != nil {
		golog.Errorf("@@@ uploadWxMedia %s", err1)
		return ""
	}
	if resultJson.Get("errcode").MustInt() != 0 {
		return ""
	}
	return resultJson.Get("mediaId").MustString()
}

func SendGzhMsg(tpl_id,touser string,data map[string]interface{}) string {
	var (
		url ="https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+getAccessTokenGzh()
		json = map[string]interface{}{}
	)
	json["touser"]=touser
	json["template_id"]=tpl_id
	json["url"]="http://weixin.qq.com/download"
	json["data"]=data
	json["miniprogram"]= map[string]string{
		"appid":wechatxcxConf().Appid,
		"pagepath":"pages/index/index",
		//"pagepath":"/userPages/whaleBean/newsActivity",

	}
	/*json = map[string]interface{} {
		"touser":touser,
		"template_id" : tpl_id,
		"url" : "http://weixin.qq.com/download",
		"data":data,
		"miniprogram":[...]interface{}{
			map[string]interface{}{
				"appid":xcx1.Appid,
				"pagepath":"pages/index/index",
			},
		},
	}*/
	if len(json) == 0 {
		golog.Errorf("wx: %s is not found")
		return ""
	}
	h := utils.NewHttpSend(url)
	h.SetSendType("JSON")
	h.SetBody(json)
	strb, err := h.Post()
	if err != nil {
		golog.Errorf("wx: sendGzhMsg http Post err: %s", err.Error())
		return ""
	}
	fmt.Printf(string(strb))
	resultJson, err1 := simplejson.NewJson(strb)
	if err1 != nil {
		golog.Errorf("@@@ sendGzhMsg %s", err1)
		return ""
	}
	if resultJson.Get("errcode").MustInt() != 0 {
		return ""
	}
	return resultJson.Get("errmsg").MustString()
}

func decryptWxData(appName,sessionKey, encryptData, iv string) (map[string]interface{}, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, err
	}
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	dataBytes, err := AesDecrypt(decodeBytes, sessionKeyBytes, ivBytes)
	fmt.Println(string(dataBytes))
	m := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &m)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	temp := m["watermark"].(map[string]interface{})
	appid := temp["appid"].(string)
	if appid != wechatxcxConf().Appid {
		return nil, fmt.Errorf("invalid appid, get !%s!", appid)
	}
	return m, nil
}

func AesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	//获取的数据尾端有'/x0e'占位符,去除它
	for i, ch := range origData {
		if ch == '\x0e' {
			origData[i] = ' '
		}
	}
	//{"phoneNumber":"15082726017","purePhoneNumber":"15082726017","countryCode":"86","watermark":{"timestamp":1539657521,"appid":"wx4c6c3ed14736228c"}}//<nil>
	return origData, nil
}