package jwts

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"go-mod/core/config"
	supports2 "go-mod/gateway/supports"
	"strings"
)
type AdminUserJwt struct {
	Id        int64 `json:"id"`
	Username  string `json:"username"`
	Name  string `json:"name"`
	Rid 	 int64 `json:"rid"`
	Bid 	 int64 `json:"bid"`
	Menumap map[string]interface{} `json:"menumap"`
	Powermap map[string]interface{} `json:"powermap"`
}
type ApiUserJwt struct {
	Id      int64 `json:"id"`
	AppId	int64 `json:"app_id"`
}
type Middleware struct {
}
var adminuserjwt         *AdminUserJwt
var apiuserjwt         *ApiUserJwt

func GetAdminUserjwt() *AdminUserJwt{
	return adminuserjwt
}
func GetApiUserjwt() *ApiUserJwt{
	return apiuserjwt
}
func ServeHTTP(ctx context.Context) {
	path := ctx.Path()
	method := ctx.Method()
	// 不需要jwt验证
	if checkJwtBeforeURL(path){
		ctx.Next()
		return
	}
	if !Serve(ctx) {
		return
	}
	//后台权限控制
	if strings.Contains(path,"/admin/"){
		var (
			isParseAdminUserToken bool
			power string
		)
		adminuserjwt = nil
		if adminuserjwt, isParseAdminUserToken = ParseAdminUserToken(ctx); !isParseAdminUserToken {
			golog.Errorf("@@@ adminuser token 解析错误")
			supports2.Error(ctx, supports2.TokenParseFailur, nil)
			return
		}
		//后台主页不拦截路由
		if strings.Contains(path, "/admin/home") {
			ctx.Next()
			return
		}
		//后台权限拦截开始
		pathList := strings.Split(path, "/")
		plen := len(pathList)
		if plen >= 3 {
			admin := pathList[1]
			menu := pathList[2]
			//GET/POST/PUT/DELETE
			if admin =="admin" {
				if plen==3 {
					switch method {
					case iris.MethodGet : power = "visible"
					case iris.MethodPost : power = "add"
					case iris.MethodPut : power = "edit"
					}
				}else {
					if method == iris.MethodDelete {
						if strings.Contains(pathList[3], ",") {
							power = "batchdel"
						}else{
							power = "del"
						}
					}else if method == iris.MethodPost{
						power = pathList[3]
					}else if method == iris.MethodPut{
						power = "edit"
					}
				}
				menumap := adminuserjwt.Menumap
				powermap := adminuserjwt.Powermap
				mPvalue, _ := menumap[menu].(float64)
				pPvalue, _ := powermap[power].(float64)
				if (int64(mPvalue) & int64(pPvalue)) == int64(pPvalue){
					ctx.Next()
					return
				}
			}
		}
		golog.Errorf("@@@ 权限不足非法操作")
		supports2.Error(ctx, supports2.PermissionsLess, nil)
		return
	}
	// Pass to real API
	if strings.Contains(path,"/api/"){
		var (
			isParseApiUserToken bool
		)
		apiuserjwt = nil
		if apiuserjwt, isParseApiUserToken = ParseApiUserToken(ctx); !isParseApiUserToken {
			golog.Errorf("@@@ apiuser token 解析错误")
			supports2.Error(ctx, supports2.TokenParseFailur, nil)
			return
		}
	}
	ctx.Next()
}

/**
return
	true:则跳过不需验证，如登录接口等...
	false:需要进一步验证
*/
func checkJwtBeforeURL(reqPath string) bool {
	for _, v := range config.O.IgnoreURLs {
		if strings.Contains(reqPath, v){
			return true
		}
	}
	return false
}
