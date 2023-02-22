package src

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/0xunion/exercise_back/src/routers"
)

func Start() {
	beego.Run()
}
