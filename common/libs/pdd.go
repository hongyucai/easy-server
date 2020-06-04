package libs

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/golog"
	"go-xm/inits"
	"go-xm/inits/parse"
	"go-xm/utils"
	"reflect"
	"sort"
	"strings"
)

/*func isPrimaryType(param interface{}) bool {
	switch param.(type) {
	case string:
		return true
	case int:
		return true
	case int64:
		return true
	case float32:
		return true
	case float64:
		return true
	case bool:
		return true
	default:
		return false
	}
}*/
func uniqueParams(params map[string]interface{}) map[string]interface{} {
	data :=map[string]interface{}{}
	for k, v := range params {
		if reflect.TypeOf(v).String() == "[]interface{}" || reflect.TypeOf(v).Kind() == reflect.Map{
			vjson,_ :=json.Marshal(v)
			data[k] = string(vjson)
		}else {
			data[k] =fmt.Sprintf("%v",v)
		}
	}
	return data
}
func PddSign(params map[string]interface{},clientSecret string) (str string) {
	var lst []string
	for k := range params {
		lst = append(lst, k)
	}
	sort.Strings(lst)
	str = ""
	for _, v := range lst {
		if v !="sign" && params[v] !="" && reflect.TypeOf(params[v]).String() != "[]interface{}" && reflect.TypeOf(params[v]).Kind() != reflect.Map {
			str += fmt.Sprintf("%s%s",v,params[v])
		}
	}
	str = strings.Trim(str,"&")
	str = clientSecret + str + clientSecret
	str = utils.Md5(str)
	str = strings.ToUpper(str) // 小写ToLower 大写ToUpper
	return
}

func syncInvoke(method string,params map[string]interface{})[]byte{
	var pdd parse.PddConfigInfo
	if inits.PddPro {
		pdd = parse.PddConfig.Pddpro
	}else {
		pdd = parse.PddConfig.Pdddev
	}
	params["client_id"]=pdd.ClientId
	//if pdd.AccessToken != ""{
	//	params["access_token"]=pdd.AccessToken
	//}
	params["version"]="V1"
	params["data_type"]="JSON"
	params["type"]=method
	params["timestamp"]=utils.GetNowUnix()

	data := uniqueParams(params)
	data["sign"]= PddSign(data,pdd.ClientSecret)
	//fmt.Println(data)
	//POST
	h := utils.NewHttpSend(pdd.ApiServerUrl)
	h.SetBody(data)
	strb, err := h.Post()
	////GET
	//h := utils.NewHttpSend(utils.GetUrlBuild(pdd.ApiServerUrl,data))
	//strb, err := h.Get()

	if err != nil {
		panic(err.Error())
	}
	//strs := strings.Replace(string(strb),"\\", "",-1)
	//fmt.Printf(strs)
	//return []byte(strs)
	return strb
}

func GetPddCategoryList(paramMap map[string]interface{}) map[string]interface{} {
	method :="pdd.goods.cats.get"
	//paramMap = map[string]interface{}{"parent_cat_id":0}//值=0时为顶点cat_id,通过树顶级节点获取cat树
	jsonByte := syncInvoke(method,paramMap)
	resultJson, err := simplejson.NewJson(jsonByte)
	if err != nil {
		golog.Errorf("@@@ 解释失败000 %s", err)
		return nil
	}
	//fmt.Println(resultJson)
	//os.Exit(1)
	resultData := resultJson.Get("goods_cats_get_response")
	data := resultData.MustMap()
	fmt.Println(data["goods_cats_list"])
	return data
}

func GetPddCouponList(paramMap map[string]interface{}) map[string]interface{} {
	method :="pdd.ddk.coupon.info.query"
	//paramMap = map[string]interface{}{"coupon_ids":[]string{},"mall_ids":[]int64{}}//coupon_ids 优惠券id 必填 mall_ids 店铺id  非必填
	jsonByte := syncInvoke(method,paramMap)
	resultJson, err := simplejson.NewJson(jsonByte)
	if err != nil {
		golog.Errorf("@@@ 解释失败000 %s", err)
		return nil
	}
	//fmt.Println(resultJson)
	//os.Exit(1)
	resultData := resultJson.Get("ddk_coupon_info_query_response")
	data := resultData.MustMap()
	fmt.Println(data["list"])
	return data
}

func GetPddPromotionBySubunionid(paramMap map[string]interface{}) map[string]interface{} {
	method :="pdd.ddk.goods.promotion.url.generate"
	//paramMap = map[string]interface{}{"goods_id_list":[]int64{}}//goods_id_list 必填
	jsonByte := syncInvoke(method,paramMap)
	resultJson, err := simplejson.NewJson(jsonByte)
	if err != nil {
		golog.Errorf("@@@ 解释失败000 %s", err)
		return nil
	}
	//fmt.Println(resultJson)
	//os.Exit(1)
	resultData := resultJson.Get("goods_promotion_url_generate_response")
	data := resultData.MustMap()
	fmt.Println(data["goods_promotion_url_list"])
	return data
}

func GetPddOrderList(paramMap map[string]interface{}) (map[string]interface{}) {
	method :="pdd.ddk.order.list.range.get"
	jsonByte := syncInvoke(method,paramMap)
	resultJson, err := simplejson.NewJson(jsonByte)
	if err != nil {
		golog.Errorf("@@@ 解释失败000 %s", err)
		return nil
	}
	//fmt.Println(resultJson)
	//os.Exit(1)
	resultData := resultJson.Get("order_list_get_response")
	data := resultData.MustMap()
	//fmt.Println(data["order_list"])
	return data
}

func GetPddGoodsList(paramMap map[string]interface{}) (map[string]interface{}) {
	method :="pdd.ddk.goods.search"
	jsonByte := syncInvoke(method,paramMap)
	resultJson, err := simplejson.NewJson(jsonByte)
	if err != nil {
		golog.Errorf("@@@ 解释失败000 %s", err)
		return nil
	}
	//fmt.Println(resultJson)
	//os.Exit(1)
	resultData := resultJson.Get("goods_search_response")
	data := resultData.MustMap()
	//fmt.Println(data["goods_list"])
	return data
}