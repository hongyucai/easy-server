package libs

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/golog"
	"go-xm/inits/parse"
	"go-xm/utils"
	"go-xm/utils/storage/redis"
	"math/rand"
	"os"
)

type paramFormt struct {
	Json       string `json:"json"`
	HttpErrors string `json:"http_errors"`
}
var yiyun = parse.YiyunConfig.Yiyundev //开发配置
//var yiyun =parse.YiyunConfig.Yiyunpro //生产配置

func sendPost(url string,paramMaps map[string]interface{},field string)  string {
	//post
	golog.Info("yiyun url:",url)
	//j,e := json.Marshal(paramMaps)
	//if e!=nil{
	//	golog.Errorf("yiyun: paramerr: %s", e.Error())
	//}
	//golog.Info("yiyun paramMaps:",paramMaps)
	h := utils.NewHttpSend( url )
	h.SetSendType("JSON" )
	h.SetBody(paramMaps)
	str, err := h.Post()

	if err != nil {
		golog.Errorf("yiyun: login http Post err: %s", err.Error())
		return ""
	}
	os.Exit(1)
	resultJson, err := simplejson.NewJson(str)
	golog.Info( "@@@ result",resultJson )
	resultStr,err := resultJson.Get(field).String()
	if err != nil{
		golog.Errorf( "@@@ 解释失败 %s",err.Error() )
		resultStr,err = resultJson.Get("vcResult").String()
		golog.Info( "@@@ 返回结果 %s",resultStr )
		return resultStr
	}
	return resultStr
}

func getToken() string {
	key := "YiYunGetToken"
	if redis.RedisDB().Get(key).Val() != ""{
		return redis.RedisDB().Get(key).Val()
	}

	paramMaps := map[string]interface{} {
		"merchant":yiyun.Merchant,
		"secret" : yiyun.Secret,
	}

	var token = sendPost(yiyun.AuthUrl+"/api/oauth/get_token",paramMaps,"token")
	if token !=""{
		redis.RedisDB().Set( key,token,3600 )
	}
	return token
}

func robotUserLogin(vcRobotSerialNo string)  bool {
	url := "/api/Robot/UserLogin?vcToken="

	paramMaps := map[string]interface{} {
			"vcMerchantNo":yiyun.Merchant,
			"nAuthorize" : yiyun.Secret,
			"vcRobotSerialNo":vcRobotSerialNo,
			"nRegionCode":"440100",
			"vcCity":"广州市",
	}
	var resultStr = sendPost( yiyun.MerchantUrl + url,paramMaps,"vcResult" )
	if resultStr=="SUCCESS"{
		return true
	}
	return false
}

func RobotUserLoginOut(vcRobotSerialNo string)  bool{
	url := "/api/Robot/UserLoginOut?vcToken="

	paramMaps := map[string]interface{} {
			"vcMerchantNo":yiyun.Merchant,
			"vcRobotSerialNo" : vcRobotSerialNo,
	}
	var resultStr = sendPost( yiyun.MerchantUrl + url,paramMaps,"vcResult" )
	if resultStr=="SUCCESS"{
		return true
	}
	return false
}

func FriendGetFriendList(vcRobotSerialNo string) string {
	url := "/api/Robot/UserLoginOut?vcToken="

	paramMaps := map[string]interface{} {
			"vcMerchantNo":yiyun.Merchant,
			"vcRobotSerialNo" : vcRobotSerialNo,
	}
	var resultStr = sendPost( yiyun.MerchantUrl + url,paramMaps,"vcResult" )
	if resultStr=="SUCCESS"{
		return ""
	}
	return resultStr

}

func RobotMerchantRobotList(vcRobotSerialNos string) interface{} {
	url := "/api/Robot/MerchantRobotList?vcToken="

	paramMaps := map[string]interface{} {
			"vcMerchantNo":yiyun.Merchant,
			"vcRobotSerialNo" : vcRobotSerialNos,
	}
	var resultStr = sendPost( yiyun.MerchantUrl + url,paramMaps,"vcResult" )
	if resultStr=="SUCCESS"{
		return true
	}
	return resultStr

}

func ChatRoomGetChatRoomList(vcRobotSerialNo string,vcChatRoomSerialNo string,isOpenMessage string) interface{}  {
	url := "/api/Robot/MerchantRobotList?vcToken="

	paramMaps := map[string]interface{} {
			"vcMerchantNo":yiyun.Merchant,
			"vcRobotSerialNo" : vcRobotSerialNo,
			"vcChatRoomSerialNo" : vcChatRoomSerialNo , //群编号，不传的话，查询全部
			"isOpenMessage" : isOpenMessage ,			//是否已开通（10 是 11 否）, 0 查询全部
	}
	var resultStr = sendPost( yiyun.MerchantUrl + url,paramMaps,"vcResult" )
	if resultStr=="SUCCESS"{
		return true
	}
	return resultStr
}

func RobotModifyProfileWhatsUp(vcRobotSerialNo string,vcWhatsUp string)  interface{} {
	url := "/api/Robot/ModifyProfileWhatsUp?vcToken="

	paramMaps := map[string]interface{} {
			"vcMerchantNo":yiyun.Merchant,
			"vcRobotSerialNo" : vcRobotSerialNo,
			"vcWhatsUp" : vcWhatsUp , //个性签名
	}
	var resultStr = sendPost( yiyun.MerchantUrl + url,paramMaps,"vcResult" )
	if resultStr=="SUCCESS"{
		return true
	}
	return false
}
//个性签名
func RobotModifyProfileName(vcRobotSerialNo string,vcNickName string) bool {
	url := "/api/Robot/ModifyProfileName?vcToken="

	paramMaps := map[string]interface{} {
			"vcMerchantNo":yiyun.Merchant,
			"vcRobotSerialNo" : vcRobotSerialNo,
			"vcWhatsUp" : vcNickName , //个性签名
	}
	var resultStr = sendPost( yiyun.MerchantUrl + url,paramMaps,"vcResult" )
	if resultStr=="SUCCESS"{
		return true
	}
	return false

}

func ChatRoomModifyGroupName(vcRobotSerialNo string,vcNickName string,vcChatRoomName string) bool {
	url := "/api/Robot/ModifyProfileName?vcToken="

	paramMaps := map[string]interface{} {
			"vcMerchantNo":yiyun.Merchant,
			"vcRobotSerialNo" : vcRobotSerialNo,
			"vcChatRoomSerialNo" : vcNickName ,
			"vcChatRoomName":vcChatRoomName,
	}
	var resultStr = sendPost( yiyun.MerchantUrl + url,paramMaps,"vcResult" )
	if resultStr=="SUCCESS"{
		return true
	}
	return false
}

func ChatRoomSetAliasForGroup(vcRobotSerialNo string,vcChatRoomSerialNo string,vcAlias string) bool {
	url := "/api/Robot/ModifyProfileName?vcToken="

	paramMaps := map[string]interface{} {
			"vcMerchantNo":yiyun.Merchant,
			"vcRobotSerialNo" : vcRobotSerialNo,
			"vcChatRoomSerialNo" : vcChatRoomSerialNo ,
			"vcAlias":vcAlias,
	}
	var resultStr = sendPost( yiyun.MerchantUrl + url,paramMaps,"vcResult" )
	if resultStr=="SUCCESS"{
		return true
	}
	return false
}

func ChatRoomGetChatRoomUserInfo(vcRobotSerialNo string,vcChatRoomSerialNo string) interface{} {

	url := "/api/ChatRoom/GetChatRoomUserInfo?vcToken="

	paramMaps := map[string]interface{} {
			"vcMerchantNo":yiyun.Merchant,
			"vcRobotSerialNo" : vcRobotSerialNo,
			"vcChatRoomSerialNo" : vcChatRoomSerialNo ,
	}
	var resultStr = sendPost( yiyun.MerchantUrl + url,paramMaps,"vcResult" )
	if resultStr=="SUCCESS"{
		return true
	}
	return false
}

func ChatRoomRobotChatRoomOpen(vcRobotSerialNo string,vcChatRoomSerialNo string) bool {
	url := "/api/ChatRoom/RobotChatRoomOpen?vcToken="

	paramMaps := map[string]interface{} {
			"vcMerchantNo":yiyun.Merchant,
			"vcRobotSerialNo" : vcRobotSerialNo,
			"vcChatRoomSerialNo" : vcChatRoomSerialNo ,

	}
	var resultStr = sendPost( yiyun.MerchantUrl + url,paramMaps,"vcResult" )
	if resultStr=="SUCCESS"{
		return true
	}
	return false
}

func ChatMessagesSendGroupChatMessages(vcRobotSerialNo string,vcChatRoomSerialNo string,nMsgType int64,msgContent string) bool {

	url := yiyun.MerchantUrl + "/api/ChatMessages/SendGroupChatMessages?vcToken=" + getToken()

	paramMaps := map[string]interface{} {
			"vcMerchantNo":yiyun.Merchant,
			"vcRobotSerialNo" : vcRobotSerialNo,
			"vcChatRoomSerialNo" : vcChatRoomSerialNo ,
			"vcRelaSerialNo": string( rand.Int() ),
			"nIsHit":1, //是否艾特 (0 艾特群内所有人 1 艾特或者不艾特用户)
			"vcToWxSerialNo":"",
			"Data":[...]interface{}{
				map[string]interface{}{
				"nMsgNum":    1,          //消息编号(整型,用于区分同一组的消息)
				"nMsgType":   nMsgType,   //消息类型 2001 文字 2002 图片 2003 语音(只支持amr格式) 2004 视频 2005 链接 2006 好友名片 2010 文件 2013小程序 2016 音乐
				"msgContent": msgContent, //文字内容（如果是图片或者链接则传图片地址[链接的图片不宜过大，建议160x160px，小于10k];如果是好友名片，则传好友的微信编号；如果是语音,则传语音的地址（语音的后缀必须为amr示例：http://downsc.chinaz.net/Files/DownLoadsound1/201910/12087.amr）;如果是小程序则传小程序的XML；如果是视频消息，则传视频的封面图【视频第一帧的图片链接地址，此处必传，否则视频消息类型会失败】）
				"nVoiceTime": 0,          //语音时长/视频时长,时长单位：秒；当消息类型为以上两种类型时，必须传时长且时长要正确，否则会发送失败，当时长不正确时可能会有很大的禁封风险
				"vcHref":     "",         //链接URL，当消息为视频时，此处传视频的链接地址
				"vcTitle":    "",         //链接标题
				"vcDesc":     "",         //链接描述
				},
			},
	}
	// post
	var resultStr = sendPost(  url,paramMaps,"vcResult" )
	fmt.Println(resultStr)
	if resultStr=="SUCCESS"{
		return true
	}
	return false
}

// 模拟三元操作符
func If(condition bool, whenTrue, whenFalse interface{}) interface{} {
	if condition {
		return whenTrue
	}
	return whenFalse
}