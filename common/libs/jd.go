package libs

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/golog"
	"go-xm/inits"
	"go-xm/inits/parse"
	"go-xm/utils"
	"sort"
	"strings"
)
//https://union.jd.com/openplatform/api
func JdSign(params map[string]string,appSecret string) (str string) {
	var lst []string
	for k := range params {
		lst = append(lst, k)
	}
	sort.Strings(lst)
	str = ""
	for _, v := range lst {
		if v !="sign" && params[v] !="" {
			str += fmt.Sprintf("%s%s",v,params[v])
		}
	}
	str = strings.Trim(str,"&")
	str = appSecret + str + appSecret
	str = utils.Md5(str)
	str = strings.ToUpper(str) // 小写ToLower 大写ToUpper
	return
}

func SetJdParamData(method string,paramJson string)[]byte{
	var jd parse.JdConfigInfo
	if inits.JdPro{
		jd = parse.JdConfig.Jdpro
	}else {
		jd = parse.JdConfig.Jddev
	}
	data := map[string]string{
		"app_key":     jd.AppKey,
		"format":      jd.Format,
		"method":      method,
		"param_json":  paramJson,
		"sign_method": jd.SignMethod,
		"timestamp":   utils.GetNowTime(),
		"v":           jd.V,
	}
	data["sign"]= JdSign(data,jd.AppSecret)
	////POST
	//h := utils.NewHttpSend(jd.Url)
	////h.SetHeader(map[string]string{"cs": "cs"})
	//h.SetBody(data)
	//strb, err := h.Post()

	//GET
	h := utils.NewHttpSend(utils.GetUrlBuild(jd.Url,data))
	strb, err := h.Get()

	if err != nil {
		panic(err.Error())
	}
	//strs := strings.Replace(string(strb),"\\", "",-1)
	//fmt.Printf(strs)
	//return []byte(strs)
	return strb
}

func GetCategoryList(paramMap map[string]interface{}) []interface{} {
	method :="jd.union.open.category.goods.get"
	paramMaps := map[string]interface{}{"req":paramMap}
	mjson,_ :=json.Marshal(paramMaps)
	paramJson :=string(mjson)
	jsonStr := SetJdParamData(method,paramJson)
	resultJson, err := simplejson.NewJson(jsonStr)
	if err != nil {
		golog.Errorf("@@@ 解释失败000 %s", err)
		return nil
	}
	resultStr, err1 := resultJson.Get("jd_union_open_category_goods_get_response").Get("result").String()
	//fmt.Printf(result)
	if err1 != nil {
		golog.Errorf("@@@ 解释失败111 %s", err1)
		return nil
	}
	dataJson, err2 := simplejson.NewJson([]byte(resultStr))
	if err2 != nil {
		golog.Errorf("@@@ 解释失败222 %s", err2)
		return nil
	}
	data, err3 := dataJson.Get("data").Array()
	if err3 != nil {
		golog.Info("@@@ 没有数据 333 %s", err3)
		return nil
	}

	//fmt.Printf("data type:%T\n", data)
	/*for _, row  := range data {
		if col, ok := row.(map[string]interface{}); ok {
			fmt.Printf("id=%d, name=%s, grade=%d, parentId=%d\n", col["id"],col["name"] ,col["grade"],col["parentId"])
			//类型判断
			if name, ok := col["name"].(string); ok {
				fmt.Println(name)
			}
			///可以看到col["parentId"]类型是json.Number
			//而json.Number是golang自带json库中decode.go文件中定义的: type Number string
			//因此json.Number实际上是个string类型
			fmt.Println(reflect.TypeOf(col["parentId"]))
			if id, ok := col["parentId"].(json.Number); ok {
				//string(n) //json.Number 类型转string
				//strconv.ParseFloat(string(n), 64) //json.Number 类型转float64
				//strconv.ParseInt(string(n), 10, 64) //json.Number 类型转int64
				id_int, err := strconv.ParseInt(string(id), 10, 64)
				if err == nil {
					fmt.Println(id_int)
				}
			}

		}
	}*/
	return data
	//return nil
}

func GetCouponList(paramMap [1]string) []interface{} {
	method :="jd.union.open.coupon.query"
	paramMaps := map[string]interface{}{"couponUrls":paramMap}
	mjson,_ :=json.Marshal(paramMaps)
	paramJson :=string(mjson)
	jsonStr := SetJdParamData(method,paramJson)
	resultJson, err := simplejson.NewJson(jsonStr)
	if err != nil {
		golog.Errorf("@@@ 解释失败000 %s", err)
		return nil
	}
	//fmt.Printf(string(paramMap[0]))
	//fmt.Printf(string(jsonStr))
	resultStr, err1 := resultJson.Get("jd_union_open_coupon_query_response").Get("result").String()
	fmt.Printf(resultStr)
	if err1 != nil {
		golog.Errorf("@@@ 解释失败111 %s", err1)
		return nil
	}
	dataJson, err2 := simplejson.NewJson([]byte(resultStr))
	if err2 != nil {
		golog.Errorf("@@@ 解释失败222 %s", err2)
		return nil
	}

	data, err3 := dataJson.Get("data").Array()
	if err3 != nil {
		golog.Info("@@@ 没有数据 333 %s", err3)
		return nil
	}
	return data
}

func GetPromotionBySubunionid(paramMap map[string]interface{}) string {
	method :="jd.union.open.promotion.bysubunionid.get"
	paramMaps := map[string]interface{}{"promotionCodeReq":paramMap}
	mjson,_ :=json.Marshal(paramMaps)
	paramJson :=string(mjson)
	jsonStr := SetJdParamData(method,paramJson)
	resultJson, err := simplejson.NewJson(jsonStr)
	if err != nil {
		golog.Errorf("@@@ 解释失败000 %s", err)
		return ""
	}
	//fmt.Printf(string(jsonStr))
	resultStr, err1 := resultJson.Get("jd_union_open_promotion_bysubunionid_get_response").Get("result").String()
	//fmt.Printf(resultStr)
	if err1 != nil {
		golog.Errorf("@@@ 解释失败111 %s", err1)
		return ""
	}
	dataJson, err2 := simplejson.NewJson([]byte(resultStr))
	if err2 != nil {
		golog.Errorf("@@@ 解释失败222 %s", err2)
		return ""
	}
	//fmt.Printf( "d",dataJson.Get("data").Get("shortURL") )
	data, err3 := dataJson.Get("data").Get("shortURL").String()
	if err3 != nil {
		golog.Info("@@@ 没有数据 333 %s", err3)
		return ""
	}
	return data
}

func GetOrderList(paramMap map[string]interface{}) ([]interface{},bool) {
	method :="jd.union.open.order.query"
	paramMaps := map[string]interface{}{"orderReq":paramMap}
	mjson,_ :=json.Marshal(paramMaps)
	paramJson :=string(mjson)
	jsonStr := SetJdParamData(method,paramJson)
	resultJson, err := simplejson.NewJson(jsonStr)
	if err != nil {
		golog.Errorf("@@@ 解释失败000 %s", err)
		return nil,false
	}
	//fmt.Printf(string(jsonStr))
	resultStr, err1 := resultJson.Get("jd_union_open_order_query_response").Get("result").String()
	fmt.Printf(resultStr)
	if err1 != nil {
		golog.Errorf("@@@ 解释失败111 %s", err1)
		return nil,false
	}
	dataJson, err2 := simplejson.NewJson([]byte(resultStr))
	if err2 != nil {
		golog.Errorf("@@@ 解释失败222 %s", err2)
		return nil,false
	}

	data, err3 := dataJson.Get("data").Array()
	if err3 != nil {
		golog.Info("@@@ 没有数据 333 %s", err3)
		return nil,false
	}
	hasMore, err4 := dataJson.Get("hasMore").Bool()
	if err4 != nil {
		golog.Errorf("@@@ 解释失败44 %s", err4)
		return nil,false
	}
	return data,hasMore
}

func GetGoodsList(paramMap map[string]interface{}) ([]interface{},bool) {
	method :="jd.union.open.goods.query"
	paramMaps := map[string]interface{}{"goodsReqDTO":paramMap}
	mjson,_ :=json.Marshal(paramMaps)
	paramJson :=string(mjson)
	jsonStr := SetJdParamData(method,paramJson)
	resultJson, err := simplejson.NewJson(jsonStr)
	if err != nil {
		golog.Errorf("@@@ 解释失败000 %s", err)
		return nil,false
	}
	//fmt.Printf(string(mjson))
	//os.Exit(1)
	resultStr, err1 := resultJson.Get("jd_union_open_goods_query_response").Get("result").String()
	fmt.Printf(resultStr)
	if err1 != nil {
		golog.Errorf("@@@ 解释失败111 %s", err1)
		return nil,false
	}
	dataJson, err2 := simplejson.NewJson([]byte(resultStr))
	if err2 != nil {
		golog.Errorf("@@@ 解释失败222 %s", err2)
		return nil,false
	}

	data, err3 := dataJson.Get("data").Array()
	if err3 != nil {
		golog.Info("@@@ 没有数据 333 %s", err3)
		return nil,false
	}
	hasMore := true

	return data,hasMore
}