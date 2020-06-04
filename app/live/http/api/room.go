package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"go-mod/app/live/http/controller"
)

func ApiHub(party iris.Party) {
	var (
		api,home,test iris.Party
	)
	api = party.Party("/api")
	{
		home = api.Party("/home")
		home.Get("/user", hero.Handler(controller.IntoRoom))

		test = api.Party("/test")
		test.Get("/user/{ids:int64}", hero.Handler(controller.LiveStart))
	}
}